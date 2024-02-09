package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 示例1
	var buffer bytes.Buffer
	contents := "Simple byte buffer for marshaling data."
	fmt.Printf("Write contents %q ...\n", contents)
	buffer.WriteString(contents)
	fmt.Printf("The length of buffer: %d\n", buffer.Len())
	fmt.Printf("The capacity of buffer: %d\n", buffer.Cap())
	fmt.Println()

	// 示例2
	s := make([]byte, 7)
	n, _ := buffer.Read(s)
	fmt.Printf("%d bytes were read. (call Read)\n", n)
	fmt.Printf("The length of buffer: %d\n", buffer.Len())
	fmt.Printf("The capacity of buffer: %d\n", buffer.Cap())
}
