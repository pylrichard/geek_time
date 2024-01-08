package pipe_filter_test

// Request is the input of the filter
type Request interface {}

// Response is the output of the filter
type Response interface {}

// Filter is the definition of the data processing components
type Filter interface {
	Process(data Request) (Response, error)
}