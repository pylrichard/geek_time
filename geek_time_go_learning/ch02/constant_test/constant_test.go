package constant_test

import "testing"

/*
	The master has failed more times than the beginner has tried.
	go test -v xxx_test.go 添加-v才能输出t.Log
	VS Code中将"go.testFlags": ["-v"]加入到settings.json
*/

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestConstant(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
	t.Log(Readable, Writeable, Executable)
	//0111
	a := 7
	t.Log(a & Readable == Readable, a & Writeable == Writeable)
}