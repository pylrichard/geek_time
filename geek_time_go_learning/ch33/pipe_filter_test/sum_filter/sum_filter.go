package sum_filter

import (
	"errors"
	pf "go/geek_time/geek_time_go_learning/ch33/pipe_filter_test"
)

type SumFilter struct {
}

var SumFilterWrongFormatError = errors.New("input data should be []int")

func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

func (sf *SumFilter) Process(data pf.Request) (pf.Response, error) {
	elements, ok := data.([]int)
	if !ok {
		return nil, SumFilterWrongFormatError
	}
	ret := 0
	for _, elem := range elements {
		ret += elem
	}

	return ret, nil
}