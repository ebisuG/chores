package main

import (
	"archive/zip"
	"io"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//zipを送信する用のヘッダーを付加
	w.Header().Set("Content-Type", "applicaiton/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=hogesample.zip")

	//Creation writer for zip file
	//In this case, this writer is from http.ResponseWriter.
	//Make zip writer based on a writer for http.
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	//Create a file inside zip
	a, err := zipWriter.Create("a.txt")
	if err != nil {
		panic(err)
	}

	//Write content of file inside zip file.
	io.Copy(a, strings.NewReader("一つ目のファイルに書き込まれるテキストです"))

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
