use pnet::packet::ethernet::{EthernetPacket, MutableEthernetPacket};
use pnet::packet::ipv4::{Ipv4Packet, MutableIpv4Packet};
use pnet::packet::udp::{MutableUdpPacket, UdpPacket};
use pnet::packet::Packet;
use pnet::util::MacAddr;
use std::net::SocketAddrV4;

pub fn build_ipv4_udp_packet(
    payload: Vec<u8>,
    source: SocketAddrV4,
    dest: SocketAddrV4,
) -> Vec<u8> {
    let ethernet_size = EthernetPacket::minimum_packet_size();
    let ipv4_size = Ipv4Packet::minimum_packet_size();
    let udp_size = UdpPacket::minimum_packet_size();
    let total_size = ethernet_size + ipv4_size + udp_size + payload.len();

    let mut buffer = vec![0u8; total_size];

    // Split buffer into three mutable slices
    let (eth_buf, rest) = buffer.split_at_mut(ethernet_size);
    let (ipv4_buf, udp_buf) = rest.split_at_mut(ipv4_size + udp_size + payload.len());

    // Ethernet layer
    let mut ethernet_layer = MutableEthernetPacket::new(eth_buf).unwrap();
    ethernet_layer.set_destination(MacAddr(0xff, 0xff, 0xff, 0xff, 0xff, 0xff));
    ethernet_layer.set_source(MacAddr(0x00, 0x00, 0x00, 0x00, 0x00, 0x00));
    ethernet_layer.set_ethertype(pnet::packet::ethernet::EtherTypes::Ipv4);

    // UDP layer
    let udp_offset = ipv4_size + udp_size + payload.len();
    let mut udp_layer = MutableUdpPacket::new(&mut udp_buf[ipv4_size..udp_offset]).unwrap();
    udp_layer.set_source(source.port());
    udp_layer.set_destination(dest.port());
    udp_layer.set_length((udp_size + payload.len()) as u16);
    udp_layer.set_payload(&payload);

    // Calculate UDP checksum
    let checksum = pnet::packet::udp::ipv4_checksum(
        &udp_layer.to_immutable(),
        &source.ip().to_owned().into(),
        &dest.ip().to_owned().into(),
    );
    udp_layer.set_checksum(checksum);

    // IPv4 layer
    let mut ipv4_layer = MutableIpv4Packet::new(&mut ipv4_buf[..ipv4_size]).unwrap();
    ipv4_layer.set_version(4);
    ipv4_layer.set_header_length(5);
    ipv4_layer.set_total_length((ipv4_size + udp_layer.packet().len()) as u16);
    ipv4_layer.set_ttl(64);
    ipv4_layer.set_next_level_protocol(pnet::packet::ip::IpNextHeaderProtocols::Udp);
    ipv4_layer.set_source(*source.ip());
    ipv4_layer.set_destination(*dest.ip());
    ipv4_layer.set_payload(udp_layer.packet());

    // Calculate IPv4 checksum
    let checksum = pnet::packet::ipv4::checksum(&ipv4_layer.to_immutable());
    ipv4_layer.set_checksum(checksum);

    buffer
}
