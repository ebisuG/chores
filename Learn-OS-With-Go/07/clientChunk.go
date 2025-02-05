package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strconv"
)

func main() {
	//keep-alive 対応版
	//define messages to send

	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	//make request to send message with POST method
	request, err := http.NewRequest(
		"GET",
		"http://localhost:8888", nil,
	)

	err = request.Write(conn)
	if err != nil {
		panic(err)
	}

	//receive chunked response from server
	reader := bufio.NewReader(conn)
	response, err := http.ReadResponse(reader, request)
	if err != nil {
		//if connection is timeout, set nil to conn for re-try.
		panic(err)
	}

	//DumpResponse don't understand zipped content, set false
	dump, err := httputil.DumpResponse(response, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
	defer response.Body.Close()

	if len(response.TransferEncoding[0]) < 1 || response.TransferEncoding[0] != "chunked" {
		panic("wrong transfer coding")
	}

	for {
		//get size
		//use size to make buffer
		sizeStr, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		//parse size in hexadecimal. Close if size is zero.
		size, err := strconv.ParseInt(string(sizeStr[:len(sizeStr)-2]), 16, 64)
		if size == 0 {
			break
		}
		if err != nil {
			panic(err)
		}

		//keep buffer for the number of size and read it.
		line := make([]byte, int(size))
		reader.Read(line)
		reader.Discard(2)
		fmt.Printf("%d bytes:%s\n", size, string(line))
	}

}
