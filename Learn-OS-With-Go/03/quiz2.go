package main

import (
	"crypto/rand"
	"io"
	"os"
)

// 無限のバイト列から、1024バイト分読み込んでファイルに書き込み
func main() {
	reader := rand.Reader
	buffer := make([]byte, 1024)
	_, err := io.ReadFull(reader, buffer)
	if err != nil {
		panic(err)
	}

	randFile, err := os.Create("rand.txt")
	if err != nil {
		panic(err)
	}
	defer randFile.Close()

	randFile.Write(buffer)

}
