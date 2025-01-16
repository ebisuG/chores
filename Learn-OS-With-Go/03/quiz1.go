package main

import (
	"io"
	"os"
)

// 古いファイルの中身を新しいファイルにコピー
// インターフェースが共通なので、実装を気にせずに使える
func main() {
	oldFile, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	defer oldFile.Close()

	newFile, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	io.Copy(newFile, oldFile)
}
