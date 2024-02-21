package main

import "math"

func Num() {
	var a int = 456
	var b int = 123
	println(a + b)
	println(a - b)
	println(a * b)
	if b != 0 {
		println(a / b)
		println(a % b)
	}

	// var c float64 = 12.3
	// 编译不通过
	// println(a + c)

	// var d int32 = 12
	// 编译不通过
	// println(a + d)
}

func ExtremeNum() {
	println("float64最大值", math.MaxFloat64)
	// 没有float64最小值
	// println("float64最小值", math.MinFloat64)
	println("float64最小正数", math.SmallestNonzeroFloat64)

	println("float32最大值", math.MaxFloat32)
	// 没有float32最小值
	// println("float32最小值", math.MinFloat32)
	println("float32最小正数", math.SmallestNonzeroFloat32)

	// int和uint都有最大值最小值
}
