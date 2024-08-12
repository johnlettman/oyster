use pcap::{Capture, Linktype, Packet, Savefile};
use std::error::Error;
use std::fmt::{Debug, Formatter};
use std::sync::Mutex;

pub struct PcapSave {
    pcap_save_path: String,
    pcap_save_file: Mutex<Savefile>,
}

impl PcapSave {
    pub fn new(pcap_save_path: String) -> Result<PcapSave, Box<dyn Error>> {
        let cap = Capture::dead(Linktype::ETHERNET)?;
        let savefile = cap.savefile(pcap_save_path.clone())?;
        let pcap_save_file = Mutex::new(savefile);
        Ok(PcapSave { pcap_save_path, pcap_save_file })
    }

    pub fn write(&mut self, packet: &Packet<'_>) -> Result<(), Box<dyn Error + '_>> {
        let mut pcap_save_file = self.pcap_save_file.lock()?;
        pcap_save_file.write(packet);
        Ok(())
    }
}

impl Debug for PcapSave {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        static NATIVE: &str = "(native)";
        f.debug_struct("PcapSave")
            .field("pcap_save_path", &self.pcap_save_path)
            .field("pcap_save_file", &NATIVE)
            .finish()
    }
}

impl Clone for PcapSave {
    fn clone(&self) -> Self {
        let pcap_save_path = self.pcap_save_path.clone();
        Self::new(pcap_save_path).unwrap()
    }

    fn clone_from(&mut self, source: &Self) {
        let pcap_save_path = source.pcap_save_path.clone();
        let cap = Capture::dead(Linktype::ETHERNET).expect("asdf");
        self.pcap_save_file = Mutex::new(cap.savefile(pcap_save_path.clone()).expect("asdf"));
        self.pcap_save_path = pcap_save_path;
    }
}
