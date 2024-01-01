package type_test

import (
	"math"
	"runtime"
	"strconv"
	"testing"
)

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	var c MyInt

	//不允许隐式类型转换
	//b = a
	b = int64(a)

	//别名类型不允许隐式类型转换
	//c = b
	c = MyInt(b)

	t.Log(a, b, c)
}

func TestInt(t *testing.T) {
	cpu := runtime.GOARCH
	size := strconv.IntSize

	t.Log(cpu, size)
}

func TestString(t *testing.T) {
	var s string
	//string是值类型，默认初始化值是空字符串，不是nil
	//加上*识别空字符串
	t.Log("*" + s + "*")
	t.Log(len(s))

	//字符串判空
	if s == "" {}
	t.Logf("%T %q", s, s)
}

func TestConvert(t *testing.T) {
	i := int64(10)
	j := int8(i)
	t.Log(j)

	a := int64(math.MaxInt64)
	b := int32(a)
	t.Log(a, b)
}

func TestPtr(t *testing.T) {
	a := 1
	ptr := &a
	//不支持指针运算
	//ptr = ptr + 1

	t.Log(a, ptr)
	t.Logf("%T %T", a, ptr)
}