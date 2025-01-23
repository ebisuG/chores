package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

// func main() {
// 	reader := strings.NewReader("Example of io.SectionReader\n")
// 	//strings.Reader や bytes.Reader でラップしてから一部のみ読み込み
// 	sectionReader := io.NewSectionReader(reader, 14, 7)
// 	io.Copy(os.Stdout, sectionReader)
// }

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
	if bytes.Equal(buffer, []byte("tEXt")) {
		rawText := make([]byte, length)
		chunk.Read(rawText)
		fmt.Println(string(rawText))
	}
}

func textChunk(text string) io.Reader {
	byteData := []byte(text)
	var buffer bytes.Buffer
	//長さの書き込み
	binary.Write(&buffer, binary.BigEndian, int32(len(byteData)))
	//チャンク名の書き込み
	buffer.WriteString("tEXt")
	//秘密のデータの書き込み
	buffer.Write(byteData)
	//calculate CRC
	crc := crc32.NewIEEE()
	io.WriteString(crc, "tEXt")
	binary.Write(&buffer, binary.BigEndian, crc.Sum32())
	return &buffer
}

func readChunks(file *os.File) []io.Reader {
	//array to store chunk
	var chunks []io.Reader
	//Skip first 8 bytes
	file.Seek(8, 0)
	var offset int64 = 8

	for {
		var length int32
		err := binary.Read(file, binary.BigEndian, &length)
		if err == io.EOF {
			break
		}
		//NewSectionReader returns sectionReader with EOF after n bytes.
		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12))

		//Go to the head of next chunk
		//Current position is the end of reading
		//Go to ahead of chunk name(4 bytes) + data length + CRC(4 bytes)
		offset, _ = file.Seek(int64(length+8), 1)
	}
	return chunks
}

func main() {
	file, err := os.Open("Lenna2.png")
	if err != nil {
		panic(err)
	}

	//Put secret string into png file
	// newFile, err := os.Create("Lenna2.png")
	// if err != nil {
	// 	panic(err)
	// }
	// defer newFile.Close()

	chunks := readChunks(file)

	// //write signature
	// io.WriteString(newFile, "\x89PNG\r\n\x1a\n")
	// // 先頭に必要なIHDRチャンクを書き込み
	// io.Copy(newFile, chunks[0])
	// // テキストチャンクを追加
	// io.Copy(newFile, textChunk("ASCII PROGRAMMING++"))

	for _, chunk := range chunks[1:] {
		dumpChunk(chunk)
		// io.Copy(newFile, chunk)
	}
}
