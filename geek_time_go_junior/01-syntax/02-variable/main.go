package main

// 首字母大写，全局可用，不建议声明全局变量
var Global = "全局变量"

// 首字母小写，本包可用，子包不能用
var internal = "包内私有变量"

var (
	First  string = "1"
	second int    = 2
)

func main() {
	// 覆盖包私有变量，变成局部变量
	var second = 5
	println(second)

	var a int = 123
	println(a)

	// Go可以类型推断，整数默认是int，可省略int，浮点默认是float64
	var b = 234
	// 局部变量声明未使用会报错
	println(b)

	// uint不可省略，否则会被解释为int
	var c uint = 456
	println(c)

	// 无法通过编译，Go是强类型语言，不会自动转换
	// println(a == c)

	var (
		d string = "dd"
		e int    = 123
	)
	println(d, e)

	// 只能用于局部变量
	f := 123
	println(f)
}
