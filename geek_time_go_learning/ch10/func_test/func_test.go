package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type IntConv func (op int) int

func retMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func calculateTime(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())

		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)

	return op
}

func TestRetValue(t *testing.T) {
	ret, _ := retMultiValues()
	t.Log(ret)
}

func TestTime(t *testing.T) {
	fn := calculateTime(slowFunc)
	t.Log(fn(10))
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}

	return ret
}

func Clear() {
	fmt.Println("Clear resources")
}

func TestSum(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
}