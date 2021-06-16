package main

import (
  "fmt"
  "net"
  "bufio"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func handler(packet gopacket.Packet, socket *bufio.Writer) {
  // do something with the packet
  fmt.Println(packet.Data())
  defer socket.Flush()
  if nu, err := socket.Write(packet.Data()); err != nil {
    panic(err)
  } else {
    fmt.Println(nu)
  }
  //fmt.Println(packet);
}

//var socket net.Conn
func main() {

  var socket *bufio.Writer
  // create a wire handler
  if handle, err := pcap.OpenLive("en0", 1600, true, pcap.BlockForever); err != nil {
  panic(err)
  // set bpf filter for handler
  } else if err := handle.SetBPFFilter("port 5060"); err != nil {  // optional
    panic(err)
  } else {

    if conn, err := net.Dial("udp", "138.201.113.114:5060"); err != nil {
      panic(err)
    } else {
      socket = bufio.NewWriter(conn);
    }
    // create a packet source and handle each packet
    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
    for packet := range packetSource.Packets() {
      handler(packet, socket);
    }
  }
}
