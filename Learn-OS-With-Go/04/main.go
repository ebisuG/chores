package main

import (
	"encoding/binary"
	"fmt"
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
	file, err := os.Open("Lenna.png")
	if err != nil {
		panic(err)
	}
	chunks := readChunks(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}
