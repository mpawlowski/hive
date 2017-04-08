package main

import (
	"flag"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"go.uber.org/zap"
)

/*
NewLogger create a logger
*/
func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"stdout",
	}
	return cfg.Build()
}

func main() {

	iface := flag.String("i", "eth0", "network interface to capture")
	flag.Parse()

	logger, _ := NewLogger()

	logger.Info("core",
		zap.String("status", "started"),
	)

	if handle, err := pcap.OpenLive(*iface, 1600, true, pcap.BlockForever); err != nil {
		logger.Error("core",
			zap.String("status", "err"),
			zap.String("msg", err.Error()),
		)
		panic(err)
	} else if err := handle.SetBPFFilter("ip"); err != nil {
		logger.Error("core",
			zap.String("status", "err"),
			zap.String("msg", err.Error()),
		)
		panic(err)
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			logger.Info("packet",
				zap.String("src.ip", packet.NetworkLayer().NetworkFlow().Src().String()),
				zap.String("src.port", packet.TransportLayer().TransportFlow().Src().String()),
				zap.String("src.mac", packet.LinkLayer().LinkFlow().Src().String()),
				zap.String("dst.ip", packet.NetworkLayer().NetworkFlow().Dst().String()),
				zap.String("dst.port", packet.TransportLayer().TransportFlow().Dst().String()),
				zap.String("dst.mac", packet.LinkLayer().LinkFlow().Dst().String()),
				zap.Int("bytes", len(packet.Data())),
			)
		}
	}

}
