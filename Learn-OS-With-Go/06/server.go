package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("Serve is running at 8080")

	//処理を一度で終了させない。常時リクエストを受付可能にする。
	for {
		//ソケット作成
		conn, err := listener.Accept()
		if err != nil {
			//handle error
			panic(err)
		}

		//非同期処理
		// go func() {
		// 	fmt.Printf("Accept %v\n", conn.RemoteAddr())
		// 	//use conn to write and read
		// 	request, err := http.ReadRequest(bufio.NewReader(conn))

		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	//io.Readerからバイト列を読み込んで分析してデバッグ出力に出す
		// 	dump, err := httputil.DumpRequest(request, true)
		// 	if err != nil {
		// 		panic(err)
		// 	}

		// 	fmt.Println(string(dump))

		// 	//http.Response構造体はWrite()メソッドを持っているので、作成したレスポンスのコンテンツをio.Writerに直接書き込む
		// 	response := http.Response{
		// 		StatusCode: 200,
		// 		ProtoMajor: 1,
		// 		ProtoMinor: 0,
		// 		Body:       ioutil.NopCloser(strings.NewReader("Hello world\n")),
		// 	}
		// 	response.Write(conn)
		// 	conn.Close()
		// }()

		//keep-alive 対応版
		go func() {
			fmt.Printf("Accept %v\n", conn.RemoteAddr())
			//最初に作成したconnを再利用するため、for-loopで繰り返しrequestを読み込む
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
				content := "Hello World\n"
				response := http.Response{
					StatusCode:    200,
					ProtoMajor:    1,
					ProtoMinor:    1,
					ContentLength: int64(len(content)),
					Body:          io.NopCloser(strings.NewReader(content)),
				}
				response.Write(conn)
			}
			conn.Close()
		}()
	}

}
