package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {
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
			conn, err = net.Dial("tcp", "localhost:8888")
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
		request.Header.Set("Accept-encoding", "gzip")

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

		//DumpResponse don't understand zipped content, set false
		dump, err := httputil.DumpResponse(response, false)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
		defer response.Body.Close()

		if response.Header.Get("Content-Encoding") == "gzip" {
			reader, err := gzip.NewReader(response.Body)
			if err != nil {
				panic(err)
			}
			io.Copy(os.Stdout, reader)
			reader.Close()
		} else {
			io.Copy(os.Stdout, response.Body)
		}

		current++
		if current == len(sendMessage) {
			break
		}

	}

}
