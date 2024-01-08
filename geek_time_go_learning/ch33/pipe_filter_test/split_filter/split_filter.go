package split_filter

import (
	"errors"
	pf "go/geek_time/geek_time_go_learning/ch33/pipe_filter_test"
	"strings"
)

type SplitFilter struct {
	delimiter string
}

var SplitFilterWrongFormatError = errors.New("input data should be string")

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter}
}

func (sf *SplitFilter) Process(data pf.Request) (pf.Response, error) {
	str, ok := data.(string)
	if !ok {
		return nil, SplitFilterWrongFormatError
	}
	parts := strings.Split(str, sf.delimiter)

	return parts, nil
}