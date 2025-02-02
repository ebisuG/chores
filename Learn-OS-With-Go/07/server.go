package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}

	fmt.Println("Server is running at port 8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go processSession(conn)
	}
}

func isGZipAcceptable(request *http.Request) bool {
	return strings.Index(strings.Join(request.Header["Accept-Encoding"], ","), "gzip") != -1
}

func processSession(conn net.Conn) {
	fmt.Println("Accept %v\n", conn.RemoteAddr())
	defer conn.Close()

	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		request, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			neterr, ok := err.(net.Error) //ダウンキャストを行い。net.Connによるエラーか、io.Readerによるエラーか判定する
			//タイムアウト時
			if ok && neterr.Timeout() {
				fmt.Println("Timeout")
				break
				//ソケットクローズ時
			} else if err == io.EOF {
				break
			}
			panic(err)
		}
		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		response := http.Response{
			StatusCode: 200,
			ProtoMajor: 1,
			ProtoMinor: 1,
			Header:     make(http.Header),
		}

		if isGZipAcceptable(request) {
			content := "Hello world(gzipped)\n"

			var buffer bytes.Buffer
			writer := gzip.NewWriter(&buffer)
			io.WriteString(writer, content)
			writer.Close()
			//make readCloser
			//ReadCloser is an interface that has reader and closer
			response.Body = io.NopCloser(&buffer)
			response.ContentLength = int64(buffer.Len())
			response.Header.Set("Content-Encoding", "gzip")
		} else {
			content := "Hello world\n"
			//assign the result of reading content.
			//
			response.Body = io.NopCloser((strings.NewReader(content)))
			response.ContentLength = int64(len(content))
		}
		response.Write(conn)
	}

}
