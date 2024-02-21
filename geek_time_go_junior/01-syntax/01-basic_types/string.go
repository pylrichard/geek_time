package main

import (
	"fmt"
	"unicode/utf8"
)

func String() {
	// 不用手写转义，IDE可以自动转义"
	// He said:"hello, go!"
	println("He said:\"hello, go!\"")
	println(`我可以换行
这是新的行
但是这里不能有反引号
`)

	println("Hello " + "Go!")
	// 字符串只能和字符串拼接
	// println("Hello " + 123)
	println("Hello " + string(123))
	println("Hello " + fmt.Sprint(123))

	// strings包有各种字符串相关操作API

	// 输出6，字节长度和编码无关
	println(len("你好"))
	// 输出2，字符数量和编码有关，通过编码库API计算
	println(utf8.RuneCountInString("你好"))
	// 输出4
	println(utf8.RuneCountInString("你好ab"))
}
