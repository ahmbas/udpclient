package main

import (
	"fmt"
	"net"
)

func runServer(host string, port int) {

	source := net.UDPAddr{IP: net.ParseIP(host), Port: port}
	sourceConn, err := net.ListenUDP("udp", &source)

	if err != nil {
		fmt.Printf("Could not establish connection on %v %v", source, err)
	}
	fmt.Printf("Starting UDP echo server on %v:%v", host, port)

	for {
		buffer := make([]byte, 10240)

		n, addr, err := sourceConn.ReadFromUDP(buffer)

		if err != nil {
			fmt.Printf("Could not read incoming packet from %v", addr)
		}
		fmt.Println(string(buffer[:n]))

	}
}
