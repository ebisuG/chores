package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	//check arguments
	if len(os.Args) == 1 {
		fmt.Printf("%s [exec file name]", os.Args[0])
		os.Exit(1)
	}
	info, err := os.Stat(os.Args[1])
	if err == os.ErrNotExist {
		fmt.Printf("file not found: %s\n", os.Args[1])
	} else if err != nil {
		panic(err)
	}
	fmt.Println("FileInfo")
	fmt.Printf("  ファイル名: %v\n", info.Name())
	fmt.Printf("  サイズ: %v\n", info.Size())
	fmt.Printf("  変更日時 %v\n", info.ModTime())
	fmt.Println("Mode()")
	fmt.Printf("  ディレクトリ？ %v\n", info.Mode().IsDir())
	fmt.Printf("  読み書き可能な通常ファイル？ %v\n", info.Mode().IsRegular())
	fmt.Printf("  ファイルアクセス権限ビット %o\n", info.Mode().Perm())
	fmt.Printf("  モードのテキスト表現 %v\n", info.Mode().String())

	//study for pointer
	//*(type) is pointer type
	//pointer is "type". "pointer type" holds adress, it's like string type holds string.

	//info.Sys() returns *syscall.Win32FileAttributeData type. This is pointer to info.Sys().
	internalStat := info.Sys().(*syscall.Win32FileAttributeData)
	// var sample *int
	// var sample2 = 2
	// sample = &sample2
	// fmt.Println(sample)
	fmt.Printf("OS固有のファイル属性:作成時間 %v\n", &internalStat.CreationTime.HighDateTime)
	fmt.Printf("OS固有のファイル属性:最終アクセス日時 %v\n", internalStat.LastAccessTime)
}
