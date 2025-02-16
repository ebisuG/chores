package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/edsrzf/mmap-go"
)

//wrap syscall.Mmap() api

func main() {
	// テストデータを書き込み
	var testData = []byte("0123456789ABCDEF")
	var testPath = filepath.Join(os.TempDir(), "testdata")
	err := os.WriteFile(testPath, testData, 0644)
	if err != nil {
		panic(err)
	}

	// メモリにマッピング
	// mは[]byteのエイリアスなので添字アクセス可能
	f, err := os.OpenFile(testPath, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	//map file to memory
	m, err := mmap.Map(f, mmap.RDWR, 0)
	if err != nil {
		panic(err)
	}
	//delete data mapped on memory
	defer m.Unmap()

	// メモリ上のデータを修正して書き込む
	m[9] = 'X'
	//書き込み
	m.Flush()

	// 読み込んでみる
	fileData, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("original: %s\n", testData)
	fmt.Printf("mmap: %s\n", m)
	fmt.Printf("file: %s\n", fileData)
}

// $ go run mapMemory.go
// original: 0123456789ABCDEF
// mmap: 012345678XABCDEF
// file: 012345678XABCDEF
