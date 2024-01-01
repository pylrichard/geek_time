package string_test_test

import (
	"strconv"
	"strings"
	"testing"
	"unsafe"
)

func TestString(t *testing.T) {
	var s string
	//默认零值为""
	t.Log("\"" + s + "\"")

	s = "hello_test"
	t.Log(len(s))
	//string是不可变的byte slice
	//s[1] = '3'

	//存储任何二进制数据
	s = "\xE4\xBA\xBB\xFF"
	t.Log(s)
	t.Log(len(s))

	s = "中"
	//输出字节数
	t.Log(len(s))

	c := []rune(s)
	t.Log(len(c))
	t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 Unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]d", c)
	}
}

func TestStringFuncs(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))

	s = strconv.Itoa(10)
	t.Log("str:" + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}