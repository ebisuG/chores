package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
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

// use channel to handle the order of response
// chan chan is channel for channel of http.Response
func writeToConn(sessionResponses chan chan *http.Response, conn net.Conn) {
	defer conn.Close()
	// 順番にResponseのchannelを取り出す
	for sessionResponse := range sessionResponses {
		// 送信データがたまるまで待つ
		// channleのデータをresponseにひとつずつ入れる
		response := <-sessionResponse
		response.Write(conn)
		close(sessionResponse)
	}
}

// セッション内のリクエストを処理する
func handleRequest(request *http.Request,
	resultReceiver chan *http.Response) {
	dump, err := httputil.DumpRequest(request, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
	content := "Hello World\n"
	// レスポンスを書き込む
	// セッションを維持するためにKeep-Aliveでないといけない
	response := &http.Response{
		StatusCode:    200,
		ProtoMajor:    1,
		ProtoMinor:    1,
		ContentLength: int64(len(content)),
		Body:          ioutil.NopCloser(strings.NewReader(content)),
	}
	// 処理が終わったらチャネルに書き込み、
	// ブロックされていたwriteToConnの処理を再始動する
	//resultReceiver is channel for response
	resultReceiver <- response
}

// セッション1つを処理
func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	// セッション内のリクエストを順に処理するためのチャネル
	//channel for channle
	sessionResponses := make(chan chan *http.Response, 50)
	defer close(sessionResponses)
	// レスポンスを直列化してソケットに書き出す専用のゴルーチン
	//リクエストの順序通りにレスポンスを返す。
	go writeToConn(sessionResponses, conn)
	reader := bufio.NewReader(conn)
	//リクエストを順序通りに処理する
	for {
		// レスポンスを受け取ってセッションのキューに
		// 入れる
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		// リクエストを読み込む
		request, err := http.ReadRequest(reader)
		if err != nil {
			neterr, ok := err.(net.Error)
			if ok && neterr.Timeout() {
				fmt.Println("Timeout")
				break
			} else if err == io.EOF {
				break
			}
			panic(err)
		}
		sessionResponse := make(chan *http.Response)
		//response用のチャネルを「response用のチャネル用」のチャネルに入れる
		sessionResponses <- sessionResponse
		// 非同期でレスポンスを実行
		go handleRequest(request, sessionResponse)
	}
}
