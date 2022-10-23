package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const mqttFixedHeaderLength int = 2

// correct values for mqtt 5.0 proto
const (
	reserved         byte = 0   // 00000000 or x00
	connect          byte = 8   // 00010000 or x08
	connAck          byte = 32  // 00100000 or x20
	publishRetain    byte = 49  // 00110001 or x31
	publishQosBitOne byte = 50  // 00110010 or x32
	publishQosBitTwo byte = 52  // 00110101 or x34
	publishDup       byte = 56  // 00111000 or x38
	pubAck           byte = 64  // 01000000 or x40
	pubRel           byte = 99  // 01100011 or x63
	pubComp          byte = 112 // 01110000 or x70
	subscribe        byte = 130 // 10000010 or x82
	subAck           byte = 144 // 10010000 or x90
	unsubscribe      byte = 162 // 10100010 or xA2
	unsubAck         byte = 176 // 10110000 or xB0
	pingReq          byte = 192 // 11000000 or xC0
	pingResp         byte = 208 // 11010000 or xD0
	disconnect       byte = 224 // 11100000 or xE0
	auth             byte = 240 // 11110000 or xF0
)

func main() {
	host := "localhost"
	port := "8080"
	log.Printf("Starting server on %s:%s\n", host, port)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalln(err)
		return
	}

	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
			return
		}
		go connectionHandler(connection)
	}
}

func connectionHandler(connection net.Conn) {
	log.Printf("Serving %s\n", connection.RemoteAddr().String())
	for {
		header, err := bufio.NewReader(connection).Peek(mqttFixedHeaderLength)
		defer connection.Close()
		if err != nil {
			// header is 2 bytes. If not present, invalid.
			log.Println("[WARN]:", err)
			// TODO flush data from reader
			return
		}

    if !validateControlPacketType(header[0]) {
      log.Println("[ERROR]: Malformed Packet")
      return
    }

    log.Println("[INFO]: Valid Packet =", header[0])


	}
}

func validateControlPacketType(b byte) bool {
  switch b {
  case connect:
  case connAck:
  case publishRetain:
  case publishQosBitOne:
  case publishQosBitTwo:
  case publishDup:
  case pubAck:
  case pubRel:
  case pubComp:
  case subscribe:
  case subAck:
  case unsubscribe:
  case unsubAck:
  case pingReq:
  case pingResp:
  case disconnect:
  case auth:
  default: 
    return false
  }
  return true
}
