package pass_array

import (
	"testing"
)

const NumOfElements = 1000

type Content struct {
	Detail [10 * NumOfElements]int
}

func withValue(a [NumOfElements]Content) int {
	//fmt.Println(&a[2])

	return 0
}

func withRef(a *[NumOfElements]Content) int {
	//e := *a
	//fmt.Println(e)

	return 0
}

func TestPassArray(t *testing.T) {
	var a [NumOfElements]Content
	withValue(a)
	withRef(&a)
}

/*
	GODEBUG=gctrace=1 go test -bench=BenchmarkPassArrayWithValue
	go test -bench=BenchmarkPassArrayWithValue -trace trace_value.out
	go tool trace trace_value.out
 */
func BenchmarkPassArrayWithValue(b *testing.B) {
	var a [NumOfElements]Content
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withValue(a)
	}
	b.StopTimer()
}

/*
	GODEBUG=gctrace=1 go test -bench=BenchmarkPassArrayWithRef
	go test -bench=BenchmarkPassArrayWithRef -trace trace_ref.out
	go tool trace trace_ref.out
 */
func BenchmarkPassArrayWithRef(b *testing.B) {
	var a [NumOfElements]Content
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withRef(&a)
	}
	b.StopTimer()
}