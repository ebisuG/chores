// io.ReaderインターフェースのReadメソッド
// func Read(p []byte) (n int, err error)

// そのまま使うと、バッファのサイズを管理しないといけない
// 1024バイトのバッファをmakeで作る
// buffer := make([]byte, 1024)
// sizeは実際に読み込んだバイト数、errはエラー
// size, err := r.Read(buffer)

// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

//標準入力からの読み込み
// func main() {
// 	for {
// 		buffer := make([]byte, 5)
// 		//標準入力からバッファサイズごとに読み込み
// 		//Os.Stdinが、io.Readerを満たしている
// 		size, err := os.Stdin.Read(buffer)
// 		if err == io.EOF {
// 			fmt.Println("EOF")
// 			break
// 		}
// 		fmt.Printf("size=%d input='%s'\n", size, string(buffer))
// 	}
// }

//インターネット通信
//下記はhttpリクエスト発行のコードだが、io.Copyの個所で通信を標準出力に一括で読み込んでいる
// package main
//
// import (
// 　　　"io"
// 　　　"os"
// 　　　"net"
// )
//
// func main() {
// 　　　conn, err := net.Dial("tcp", "ascii.jp:80")
// 　　　if err != nil {
// 　　　　　　panic(err)
// 　　　}
// 　　　conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
// 　　　io.Copy(os.Stdout, conn)
// }
