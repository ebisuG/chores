package main

import (
	"fmt"
	"net"
)

func main() {
	//udp4でないとエラーになる
	conn, err := net.Dial("udp4", "localhost:8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("sending to server")
	_, err = conn.Write([]byte("Hello from client"))
	if err != nil {
		panic(err)
	}

	fmt.Println("Reciving from server")
	buffer := make([]byte, 1500)
	//connで受け取ったデータをbufferに読み出す
	length, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Received :%s\n", string(buffer[:length]))

}
