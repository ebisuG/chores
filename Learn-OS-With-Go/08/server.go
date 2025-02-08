package main

import (
	"fmt"
	"net"
)

//TCPだと、net.Listen()関数を呼び、返ってきたnet.Listenerインタフェースでクライアントが接続してくるのを待っていた
//Unlike net.Listener.Accept(), net.ListenPacket doesn't need to wait connection from client and immediately return UDP socket

func main() {
	fmt.Println("Server is runnnig at port 8888")
	//net.PacketConn
	conn, err := net.ListenPacket("udp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buffer := make([]byte, 1500)
	for {
		//There is no .Accept(), conn.ReadFrom returns remote address at this timing
		length, remoteAddress, err := conn.ReadFrom(buffer)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Received from %v:%v\n", remoteAddress, string(buffer[:length]))
		_, err = conn.WriteTo([]byte("Hello from server"), remoteAddress)
		if err != nil {
			panic(err)
		}
	}

}
