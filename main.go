package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
)

func main() {
	args := os.Args[1:]
	ethernetInterface := args[0]
	mtuSize, err := strconv.ParseInt(args[1], 10, 32)
	if err != nil {
		log.Fatalf("MTU size: %v", err)
		return
	}
	port, err := strconv.Atoi(args[2])
	if err != nil {
		log.Fatalf("Port number: %v", err)
		return
	}
	outputFileName := args[3]
	f, err := os.Create(outputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	pcapw := pcapgo.NewWriter(f)
	if err := pcapw.WriteFileHeader(1600, layers.LinkTypeEthernet); err != nil {
		log.Fatalf("WriteFileHeader: %v", err)
	}

	handle, err := pcap.OpenLive(ethernetInterface, int32(mtuSize), true, pcap.BlockForever)
	if err != nil {
		log.Fatalf("OpenEthernet: %v", err)
	}
	err = handle.SetBPFFilter(fmt.Sprintf("tcp and port %d", port))
	if err != nil {
		log.Fatalf("BPF Filter: %v", err)
	}

	pkgsrc := gopacket.NewPacketSource(handle, layers.LayerTypeEthernet)
	for packet := range pkgsrc.Packets() {
		if err := pcapw.WritePacket(packet.Metadata().CaptureInfo, packet.Data()); err != nil {
			log.Fatalf("pcap.WritePacket(): %v", err)
		}
	}
}
