package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// 示例1
	reader := strings.NewReader(
		"NewReader returns a new Reader reading from s. " +
			"It is similar to bytes.NewBufferString but more efficient and read-only.")
	fmt.Printf("The size of reader: %d\n", reader.Size())
	fmt.Printf("The reading index in reader: %d\n",
		reader.Size()-int64(reader.Len()))

	buf1 := make([]byte, 47)
	n, _ := reader.Read(buf1)
	fmt.Printf("%d bytes were read. (call Read)\n", n)
	fmt.Printf("The reading index in reader: %d\n",
		reader.Size()-int64(reader.Len()))
	fmt.Println()

	// 示例2
	buf2 := make([]byte, 21)
	offset1 := int64(64)
	n, _ = reader.ReadAt(buf2, offset1)
	fmt.Printf("%d bytes were read. (call ReadAt, offset: %d)\n", n, offset1)
	fmt.Printf("The reading index in reader: %d\n",
		reader.Size()-int64(reader.Len()))
	fmt.Println()

	// 示例3
	offset2 := int64(17)
	expectedIndex := reader.Size() - int64(reader.Len()) + offset2
	fmt.Printf("Seek with offset %d and whence %d ...\n", offset2, io.SeekCurrent)
	readingIndex, _ := reader.Seek(offset2, io.SeekCurrent)
	fmt.Printf("The reading index in reader: %d (returned by Seek)\n", readingIndex)
	fmt.Printf("The reading index in reader: %d (computed by me)\n", expectedIndex)

	n, _ = reader.Read(buf2)
	fmt.Printf("%d bytes were read. (call Read)\n", n)
	fmt.Printf("The reading index in reader: %d\n",
		reader.Size()-int64(reader.Len()))
}
