package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	// //make tcp connection
	// conn, err := net.Dial("tcp", "localhost:8080")
	// if err != nil {
	// 	panic(err)
	// }

	// //Send get request
	// request, err := http.NewRequest("GET", "http://localhost:8080", nil)
	// if err != nil {
	// 	panic(err)
	// }

	// //write request to conn writer
	// request.Write(conn)

	// //read response from request
	// response, err := http.ReadResponse(bufio.NewReader(conn), request)
	// if err != nil {
	// 	panic(err)
	// }

	// dump, err := httputil.DumpResponse(response, true)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(dump))

	//keep-alive 対応版
	//define messages to send
	sendMessage := []string{
		"ASCII",
		"PROGRAMING",
		"PLUS",
	}
	current := 0
	var conn net.Conn = nil
	//use for block for re-try
	for {
		var err error
		if conn == nil {
			//make socket
			conn, err = net.Dial("tcp", "localhost:8080")
			if err != nil {
				panic(err)
			}
			fmt.Printf("Access:%d\n", current)
		}
		//make request to send message with POST method
		request, err := http.NewRequest(
			"POST",
			"http://localhost:8888",
			strings.NewReader(sendMessage[current]),
		)
		if err != nil {
			panic(err)
		}
		err = request.Write(conn)
		if err != nil {
			panic(err)
		}

		//read response from server
		response, err := http.ReadResponse(bufio.NewReader(conn), request)
		if err != nil {
			//if connection is timeout, set nil to conn for re-try.
			fmt.Print("Retry")
			conn = nil
			continue
		}

		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		current++
		if current == len(sendMessage) {
			break
		}

	}

}
