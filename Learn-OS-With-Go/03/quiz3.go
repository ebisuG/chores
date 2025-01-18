package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	//内容を書き込むzipファイル
	file, err := os.Create("sample.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//zipファイルへのWriter
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	//zipファイル内にファイルを作成
	a, err := zipWriter.Create("a.txt")
	if err != nil {
		panic(err)
	}
	//読み取った内容をwriterに渡す
	//圧縮ファイルの中に直接書き込める（？）
	io.Copy(a, strings.NewReader("一つ目のファイルに書き込まれるテキストです"))

	b, err := zipWriter.Create("b.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(b, strings.NewReader("二つ目のファイルに書き込まれるテキストです"))
}
