package optmization

import "testing"

func TestCreateRequest(t *testing.T) {
	ret := createRequest()
	t.Log(ret)
}

func TestProcessRequest(t *testing.T) {
	var reqs []string
	reqs = append(reqs, createRequest())
	reps := processRequestHighPerf(reqs)
	t.Log(reps[0])
}

func BenchmarkProcessRequestLowPerf(b *testing.B) {
	var reqs []string
	reqs = append(reqs, createRequest())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = processRequestLowPerf(reqs)
	}
	b.StopTimer()
}

func BenchmarkProcessRequestHighPerf(b *testing.B) {
	var reqs []string
	reqs = append(reqs, createRequest())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = processRequestHighPerf(reqs)
	}
	b.StopTimer()
}