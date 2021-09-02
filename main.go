package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/gopacket/layers"
	"github.com/mushorg/go-dpi"
	"github.com/mushorg/go-dpi/types"
	"github.com/mushorg/go-dpi/utils"
	"log"
)

type output struct {
	SrcIP   string         `json:"src_ip"`
	DstIP   string         `json:"dst_ip"`
	SrcPort uint16         `json:"src_port"`
	DstPort uint16         `json:"dst_port"`
	Proto   string         `json:"proto"`
	Class   types.Protocol `json:"class"`
}

func (this *output) String() string {

	out, err := json.Marshal(&this)
	if err != nil {
		log.Println("Marshal Err:", err)
	}

	return string(out)
}

func main() {
	godpi.Initialize()
	defer godpi.Destroy()
	packets, err := utils.ReadDumpFile("test.pcap")
	if err != nil {
		fmt.Println(err)
	} else {
		for packet := range packets {
			var o output
			flow, _ := godpi.GetPacketFlow(packet)
			result := godpi.ClassifyFlow(flow)

			ipv4, ok := packet.NetworkLayer().(*layers.IPv4)
			if !ok {
				return // Ignore packets that aren't IPv4
			}

			if ipv4.FragOffset != 0 || (ipv4.Flags&layers.IPv4MoreFragments) != 0 {
				return // Ignore fragmented packets.
			}

			o.SrcIP = ipv4.SrcIP.String()
			o.DstIP = ipv4.DstIP.String()
			o.Proto = ipv4.Protocol.String()
			o.Class = result.Protocol

			switch t := packet.TransportLayer().(type) {
			case *layers.TCP:
				o.SrcPort = uint16(t.SrcPort)
				o.DstPort = uint16(t.DstPort)
			case *layers.UDP:
				o.SrcPort = uint16(t.SrcPort)
				o.DstPort = uint16(t.DstPort)
			}
			log.Println(o.String())
		}
	}
}
