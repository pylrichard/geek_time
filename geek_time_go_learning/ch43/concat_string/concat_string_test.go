package concat_string

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

const Numbers = 100

func BenchmarkSprintf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var s string
		for j := 0; j < Numbers; j++ {
			s = fmt.Sprintf("%v%v", s, i)
		}
	}
	b.StopTimer()
}

func BenchmarkStringBuilder(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < Numbers; j++ {
			builder.WriteString(strconv.Itoa(i))
		}
		_ = builder.String()
	}
	b.StopTimer()
}

func BenchmarkBytesBuf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for i := 0; i < Numbers; i++ {
			buf.WriteString(strconv.Itoa(i))
		}
		_ = buf.String()
	}
	b.StopTimer()
}

func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var s string
		for i := 0; i < Numbers; i++ {
			//string是不可变的，要生成新string，需要开拓新空间，拷贝数据，GC还要回收旧string
			s += strconv.Itoa(i)
		}

	}
	b.StopTimer()
}