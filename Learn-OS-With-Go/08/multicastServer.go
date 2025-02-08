package main

import (
	"fmt"
	"net"
	"time"
)

// server sends time data to client as multicast
func main() {
	fmt.Println("Start tick server at 239.0.0.1:9999")
	conn, err := net.Dial("udp", "239.0.0.1:9999")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	start := time.Now()
	wait := 5*time.Second -
		time.Nanosecond*time.Duration(
			start.UnixNano()%(10*1000*1000*1000))
	time.Sleep(wait)
	ticker := time.Tick(10 * time.Second)
	for now := range ticker {
		conn.Write([]byte(now.String()))
		fmt.Println("Tick: ", now.String())
	}
}
