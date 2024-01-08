package straight_pipline

import pf "go/geek_time/geek_time_go_learning/ch33/pipe_filter_test"

// StraightPipeline is composed of the filters, and the filters are piled as a straight line
type StraightPipeline struct {
	Name	string
	Filters	*[]pf.Filter
}

func NewStraightPipeline(name string, filters ...pf.Filter) *StraightPipeline {
	return &StraightPipeline{
		Name:		name,
		Filters:	&filters,
	}
}

// Process is to process the coming data by the pipeline
func (sp *StraightPipeline) Process(data pf.Request) (pf.Response, error) {
	var ret interface{}
	var err error
	for _, filter := range *sp.Filters {
		ret, err = filter.Process(data)
		if err != nil {
			return ret, err
		}
		data = ret
	}

	return ret, err
}