package package_callee

import "fmt"

//一个源文件可以有多个init
func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

//Square 写成square则无法被外部调用
func Square(n int) int {
	return n * n
}

func Sum(a int, b int) int {
	return a + b
}