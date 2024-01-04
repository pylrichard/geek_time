package benchmark_test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcatStringByAdd(t *testing.T) {
	assert := assert.New(t)
	elements := []string{"1", "2", "3", "4", "5"}
	ret := ""
	for _, elem := range elements {
		ret += elem
	}
	assert.Equal("12345", ret)
}

func BenchmarkConcatStringByAdd(b *testing.B) {
	elements := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elements {
			ret += elem
		}
	}
	b.StopTimer()
}

func TestConcatStringByBytesBuffer(t *testing.T) {
	assert := assert.New(t)
	var buf bytes.Buffer
	elements := []string{"1", "2", "3", "4", "5"}
	for _, elem := range elements {
		buf.WriteString(elem)
	}
	assert.Equal("12345", buf.String())
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	elements := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for _, elem := range elements {
			buf.WriteString(elem)
		}
	}
	b.StopTimer()
}