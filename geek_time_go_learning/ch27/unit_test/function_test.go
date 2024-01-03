package unit_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	go test -v -cover 显示代码覆盖率
 */
func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is %d, the expected is %d, the actual %d",
				inputs[i], expected[i], ret)
		}
	}
}

func TestError(t *testing.T) {
	fmt.Println("Start")
	t.Error("error")
	fmt.Println("End")
}

func TestFatal(t *testing.T) {
	fmt.Println("Start")
	//测试中止，后面代码不执行
	t.Fatal("fatal")
	fmt.Println("End")
}

func TestSquareWithAssert(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		assert.Equal(t, expected[i], ret)
	}
}