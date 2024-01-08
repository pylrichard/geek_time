package straight_pipline

import (
	sf "go/geek_time/geek_time_go_learning/ch33/pipe_filter_test/split_filter"
	sum "go/geek_time/geek_time_go_learning/ch33/pipe_filter_test/sum_filter"
	tif "go/geek_time/geek_time_go_learning/ch33/pipe_filter_test/to_int_filter"
	"testing"
)

func TestStraightPipeline(t *testing.T) {
	splitFilter := sf.NewSplitFilter(",")
	toIntFilter := tif.NewToIntFilter()
	sumFilter := sum.NewSumFilter()
	sp := NewStraightPipeline("sp", splitFilter, toIntFilter, sumFilter)
	ret, err := sp.Process("1,2,3")
	if err != nil {
		t.Fatal(err)
	}
	if ret != 6 {
		t.Fatalf("The expected is 6, but the actual is %d", ret)
	}
}