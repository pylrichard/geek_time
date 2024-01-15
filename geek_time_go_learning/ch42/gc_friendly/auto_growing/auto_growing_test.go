package auto_growing

import "testing"

const NumOfElements = 100000
const Times = 1000

func TestAutoGrow(t *testing.T) {
	for i := 0; i < Times; i++ {
		var s []int
		for j := 0; j < NumOfElements; j++ {
			s = append(s, j)
		}
	}
}

func TestProperInit(t *testing.T) {
	for i := 0; i < Times; i++ {
		s := make([]int, 0, NumOfElements)
		for j := 0; j < NumOfElements; j++ {
			s = append(s, j)
		}
	}
}

func TestOverSizeInit(t *testing.T) {
	for i := 0; i < Times; i++ {
		s := make([]int, 0, 2 * NumOfElements)
		for j := 0; j < NumOfElements; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkAutoGrow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s []int
		for j := 0; j < NumOfElements; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkProperInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, NumOfElements)
		for j := 0; j < NumOfElements; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkOverSizeInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, 2 * NumOfElements)
		for j := 0; j < NumOfElements; j++ {
			s = append(s, j)
		}
	}
}