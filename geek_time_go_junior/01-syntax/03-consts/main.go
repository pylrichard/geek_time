package main

// 常量支持类型推断
const External = "包外可访问"
const internal = "包内可访问"

const (
	a = 123
)

const (
	// iota方便控制常量初始化
	Init = iota
	Running
	Paused
	Stopped

	// StatusB也是10
	StatusA = 10
	StatusB
)

const (
	One = iota << 1
	Two
	Four
)

func main() {
	// 常量不能被修改
	// a = 456
	println(Init, Running, StatusA, StatusB)
	println(Two, Four)
}
