package main

import (
	"io"
	"net"
	"os"
)

//処理の抽象化・共通化
//io.Writer インタフェースの定義
//バイト列を引数に取り、（書き込んだ）数字と、あればエラーを返す
// type Writer interface {
// 	    Write(p []byte) (n int, err error)
// 	}

// ファイル出力を行う場合
// func main() {
//     file, err := os.Create("test.txt")
//     if err != nil {
//         panic(err)
//     }
//     file.Write([]byte("os.File example\n"))
//     file.Close()
// }
// file.Writeがファイル出力を行っている

// 画面出力を行う場合
// func main() {
// 	    os.Stdout.Write([]byte("os.Stdout example\n"))
// 	}
// Stdout.Writeが呼ばれている

// バッファを利用する場合
// func main() {
// 	var buffer bytes.Buffer
// 	buffer.Write([]byte("bytes.Buffer example\n"))
// 	fmt.Println(buffer.String())
// }

// net.Conn は io.Writer と io.Reader のハイブリッドなインタフェース
// Writeメソッドを利用している
func main() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
	io.Copy(os.Stdout, conn)
}

// Goでは、JavaのインタフェースやC++の親クラスと異なり、「このインタフェースを持っています」という宣言を構造体側には一切書きません。
// どの構造体がこのインタフェースを満たしているかは、コードを単純に検索するだけでは探せません。
// godoc コマンドで、あるインターフェースがどの構造体から実装されているかを解析してくれる
// godoc -http ":6060" -analysis type
