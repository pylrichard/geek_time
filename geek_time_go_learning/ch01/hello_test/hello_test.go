package main

import (
	"fmt"
	"os"
)

/*
	go build .
    ./hello_test pyl
	退出码通过echo $?查看
*/
func main() {
	if len(os.Args) > 1 {
		fmt.Println("Hello", os.Args[1])
	}

	os.Exit(-1)
}