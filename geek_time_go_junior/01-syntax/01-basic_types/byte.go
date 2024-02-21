package main

import (
	"fmt"
)

func Byte() {
	var a byte = 'a'
	// 输出a的ASCII码97
	println(a)
	// 输出'a'
	println(fmt.Sprintf("%c", a))

	var str string = "this is string"
	// 转换发生了复制
	var bs []byte = []byte(str)
	println(bs)
	// 不会修改str
	bs[0] = 'T'
	println(str)
	println(string(bs))
}
