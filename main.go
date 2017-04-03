package main

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"go.uber.org/zap"
)

func handlePacket(packet gopacket.Packet) {
	logger, _ := zap.NewProduction()

	logger.Info("packet",
		zap.String("src.ip", packet.NetworkLayer().NetworkFlow().Src().String()),
		zap.String("src.port", packet.TransportLayer().TransportFlow().Src().String()),
		zap.String("dst.ip", packet.NetworkLayer().NetworkFlow().Dst().String()),
		zap.String("dst.port", packet.TransportLayer().TransportFlow().Dst().String()),
	)

}

func main() {

	logger, _ := zap.NewProduction()

	logger.Info("core",
		zap.String("status", "started"),
	)

	if handle, err := pcap.OpenLive("eno1", 1600, true, pcap.BlockForever); err != nil {
		logger.Error("core",
			zap.String("status", "err"),
			zap.String("msg", err.Error()),
		)
		panic(err)
	} else if err := handle.SetBPFFilter("tcp and port 80"); err != nil {
		logger.Error("core",
			zap.String("status", "err"),
			zap.String("msg", err.Error()),
		)
		panic(err)
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			handlePacket(packet)
		}
	}

}
