use crate::client::pcap_save::PcapSave;
use crate::packet::{ImuPacket, LidarPacket};
use binrw::{BinRead, BinReaderExt};
use pcap::Capture;
use pnet::packet::ethernet::EthernetPacket;
use pnet::packet::ipv4::Ipv4Packet;
use pnet::packet::udp::UdpPacket;
use pnet::packet::Packet;
use rosbag::record_types::Connection;
use rosbag::{ChunkRecord, MessageRecord, RosBag};
use std::error::Error;
use std::fmt::Debug;
use std::io::Cursor;
use std::net::{Ipv4Addr, SocketAddr, SocketAddrV4, ToSocketAddrs};
use std::path::Path;
use std::sync::Arc;
use tokio::net::UdpSocket;
use tokio::runtime::Runtime;
use tokio::sync::{mpsc, Mutex};
use tokio::task;

const DEFAULT_ROS_LIDAR_TOPIC: &'static str = "/ouster/lidar_packets";
const DEFAULT_ROS_IMU_TOPIC: &'static str = "/ouster/imu_packets";

const CHANNEL_BUFFER_SIZE: usize = 32;
const PACKET_BUFFER_SIZE: usize = 1024 * 1024;

/// The virtual source IPv4 address for saving captures from non-network sources.
const VIRTUAL_SOURCE_ADDR: Ipv4Addr = Ipv4Addr::new(127, 0, 0, 1);
const VIRTUAL_DEST_ADDR: Ipv4Addr = Ipv4Addr::new(127, 0, 0, 1);

const DUMMY_LIDAR_SOURCE_SOCKET_ADDR: SocketAddrV4 = SocketAddrV4::new(VIRTUAL_SOURCE_ADDR, 3);
const DUMMY_IMU_SOURCE_SOCKET_ADDR: SocketAddrV4 = SocketAddrV4::new(VIRTUAL_SOURCE_ADDR, 3);

const DUMMY_LIDAR_DEST_SOCKET_ADDR: SocketAddrV4 = SocketAddrV4::new(VIRTUAL_DEST_ADDR, 3);
const DUMMY_IMU_DEST_SOCKET_ADDR: SocketAddrV4 = SocketAddrV4::new(VIRTUAL_DEST_ADDR, 33);

#[derive(Debug)]
pub struct Client {
    lidar_tx: Option<mpsc::Sender<Vec<u8>>>,
    lidar_rx: Option<Arc<Mutex<mpsc::Receiver<Vec<u8>>>>>,
    imu_tx: Option<mpsc::Sender<Vec<u8>>>,
    imu_rx: Option<Arc<Mutex<mpsc::Receiver<Vec<u8>>>>>,
}

impl Client {
    pub const fn new() -> Client {
        Client { lidar_tx: None, lidar_rx: None, imu_tx: None, imu_rx: None }
    }

    fn lidar_channel_or_new(&mut self) -> mpsc::Sender<Vec<u8>> {
        if let Some(ref lidar_tx) = self.lidar_tx {
            lidar_tx.clone()
        } else {
            let (lidar_tx, lidar_rx) = mpsc::channel(CHANNEL_BUFFER_SIZE);
            self.lidar_tx = Some(lidar_tx.clone());
            self.lidar_rx = Some(Arc::new(Mutex::new(lidar_rx)));
            lidar_tx
        }
    }

    fn imu_channel_or_new(&mut self) -> mpsc::Sender<Vec<u8>> {
        if let Some(ref imu_tx) = self.imu_tx {
            imu_tx.clone()
        } else {
            let (imu_tx, imu_rx) = mpsc::channel(CHANNEL_BUFFER_SIZE);
            self.imu_tx = Some(imu_tx.clone());
            self.imu_rx = Some(Arc::new(Mutex::new(imu_rx)));
            imu_tx
        }
    }

    pub async fn with_lidar_socket<A: ToSocketAddrs>(
        &mut self,
        lidar_addr: A,
    ) -> Result<&Client, Box<dyn Error>> {
        let lidar_tx = self.lidar_channel_or_new();
        let lidar_addr = lidar_addr.to_socket_addrs()?.next().ok_or("Invalid LIDAR address");
        let lidar_socket = UdpSocket::bind(lidar_addr?).await?;
        let lidar_packet_tx = lidar_tx.clone();

        task::spawn(async {
            listen_for_packets(lidar_socket, lidar_packet_tx).await;
        });

        Ok(self)
    }

    pub async fn with_imu_socket<A: ToSocketAddrs>(
        &mut self,
        imu_addr: A,
    ) -> Result<&Client, Box<dyn Error>> {
        let imu_tx = self.imu_channel_or_new();
        let imu_addr = imu_addr.to_socket_addrs()?.next().ok_or("Invalid IMU address");
        let imu_socket = UdpSocket::bind(imu_addr?).await?;
        let imu_packet_tx = imu_tx.clone();

        task::spawn(async move {
            listen_for_packets(imu_socket, imu_packet_tx).await;
        });

        Ok(self)
    }

    pub fn with_pcap(
        &mut self,
        pcap_path: &str,
        lidar_addr: SocketAddr,
        imu_addr: SocketAddr,
    ) -> Result<&Client, Box<dyn Error>> {
        let lidar_tx = self.lidar_channel_or_new();
        let imu_tx = self.imu_channel_or_new();
        let pcap_path = pcap_path.to_string();

        task::spawn(async move {
            listen_for_pcap(&pcap_path, lidar_addr, imu_addr, lidar_tx, imu_tx).await.unwrap();
        });

        Ok(self)
    }

    pub fn with_bag(
        &mut self,
        bag_path: &str,
        lidar_topic: &str,
        imu_topic: &str,
    ) -> Result<&Client, Box<dyn Error>> {
        let lidar_tx = self.lidar_channel_or_new();
        let imu_tx = self.imu_channel_or_new();
        let bag_path = bag_path.to_string();
        let lidar_topic = lidar_topic.to_string();
        let imu_topic = imu_topic.to_string();

        task::spawn(async move {
            listen_for_bag_lidar(&bag_path, &lidar_topic, lidar_tx, &imu_topic, imu_tx)
                .await
                .unwrap();
        });

        Ok(self)
    }

    pub async fn process_lidar<F>(&mut self, mut callback: F) -> Result<&Client, Box<dyn Error>>
    where
        F: FnMut(LidarPacket) + Send + 'static,
    {
        if let Some(lidar_rx) = &self.lidar_rx {
            let mut lidar_rx = lidar_rx.lock().await;
            while let Some(packet) = lidar_rx.recv().await {
                let mut stream = Cursor::new(packet);
                let lidar_packet: LidarPacket = stream.read_le().unwrap();
                callback(lidar_packet);
            }
        }

        Ok(self)
    }

    pub async fn process_imu<F>(&mut self, mut callback: F) -> Result<&Client, Box<dyn Error>>
    where
        F: FnMut(ImuPacket) + Send + 'static,
    {
        if let Some(imu_rx) = &self.imu_rx {
            let mut imu_rx = imu_rx.lock().await;
            while !imu_rx.is_closed() {
                if let Some(packet) = imu_rx.recv().await {
                    let mut stream = Cursor::new(packet);
                    let imu_packet: ImuPacket = stream.read_le().unwrap();
                    callback(imu_packet);
                }
            }
        }

        Ok(self)
    }
}

async fn listen_for_packets(socket: UdpSocket, tx: mpsc::Sender<Vec<u8>>) {
    let mut buffer: Vec<u8> = vec![0u8; PACKET_BUFFER_SIZE];

    loop {
        match socket.recv_from(&mut buffer).await {
            Ok((size, addr)) => {
                let packet = buffer[..size].to_vec();
                if tx.send(packet).await.is_err() {
                    eprintln!("Receiver dropped");
                    return;
                }
            },

            Err(e) => {
                eprintln!("Error receiving packet: {e:?}");
                return;
            },
        }
    }
}

async fn listen_for_pcap(
    pcap_path: &str,
    lidar_addr: SocketAddr,
    imu_addr: SocketAddr,
    lidar_tx: mpsc::Sender<Vec<u8>>,
    imu_tx: mpsc::Sender<Vec<u8>>,
) -> Result<(), Box<dyn Error>> {
    let mut cap = Capture::from_file(pcap_path)?;

    while let Ok(packet) = cap.next_packet() {
        let ethernet = EthernetPacket::new(packet.data).unwrap();
        if let Some(ipv4) = Ipv4Packet::new(ethernet.payload()) {
            if let Some(udp) = UdpPacket::new(ipv4.payload()) {
                let source_addr = SocketAddr::new(ipv4.get_source().into(), udp.get_source());
                let data = udp.payload().to_vec();

                if source_addr == lidar_addr {
                    if lidar_tx.send(data).await.is_err() {
                        eprintln!("LIDAR receiver dropped");
                    }
                } else if source_addr == imu_addr {
                    if imu_tx.send(data).await.is_err() {
                        eprintln!("IMU receiver dropped");
                    }
                }
            }
        }
    }

    Ok(())
}

async fn listen_for_bag_lidar(
    bag_path: &str,
    lidar_topic: &str,
    lidar_tx: mpsc::Sender<Vec<u8>>,
    imu_topic: &str,
    imu_tx: mpsc::Sender<Vec<u8>>,
) -> Result<(), Box<dyn Error>> {
    let path = Path::new(bag_path);
    let bag = RosBag::new(path)?;

    let mut lidar_topic_id = u32::MAX;
    let mut imu_topic_id = u32::MAX;

    for record in bag.chunk_records() {
        let chunk_record = record.unwrap();
        match chunk_record {
            ChunkRecord::Chunk(chunk) => {
                for message in chunk.messages() {
                    match message.unwrap() {
                        MessageRecord::Connection(connection) => {
                            let Connection { topic, id, .. } = connection;

                            if topic == lidar_topic {
                                lidar_topic_id = id;
                            } else if topic == imu_topic {
                                imu_topic_id = id;
                            }
                        },
                        MessageRecord::MessageData(data) => {
                            if data.conn_id == lidar_topic_id {
                                let data_vec = data.data.to_vec();

                                if lidar_tx.send(data_vec).await.is_err() {
                                    eprintln!("LIDAR receiver dropped");
                                }
                            }
                        },
                    }
                }
            },
            _ => {},
        }
    }

    Ok(())
}

#[tokio::test]
async fn test() {
    let lidar_addr: SocketAddr = "127.0.0.1:7502".parse().unwrap();
    let imu_addr: SocketAddr = "127.0.0.1:7503".parse().unwrap();
    let mut client = Client::new();
    client.with_pcap("/home/jlettman/repos/oyster/samples/pcaps/OS-0-32-U1_v2.2.0_1024x10-single-packet.pcap", lidar_addr, imu_addr).expect("Failed to set up PCAP");

    let process_lidar_task = tokio::spawn(async move {
        client
            .process_lidar(|packet: LidarPacket| {
                println!("Received LIDAR packet: {:?}", packet);
            })
            .await
            .unwrap();
    });

    process_lidar_task.await.unwrap();
}
