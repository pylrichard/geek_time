package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 2, 4}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}

	//元素值和顺序不同，判断为不同
	t.Log(a == b)
	//数组维数不同不能比较
	//t.Log(a == c)
	t.Log(a == d)
}

func TestBitClear(t *testing.T) {
	//0111
	a := 7
	//对应位置清零
	a = a &^ Readable
	a = a &^ Executable
	t.Log(a & Readable == Readable,
		a & Writeable == Writeable,
		a & Executable == Executable)
}