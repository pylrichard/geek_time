package optmization

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

/*
	go test -bench=. -cpuprofile=cpu.prof -memprofile=mem.prof
	go tool pprof cpu.prof
	go tool pprof mem.prof
	list processRequestHighPerf
	list processRequestLowPerf
 */
func createRequest() string {
	payload := make([]int, 100, 100)
	for i := 0; i < 100; i++ {
		payload[i] = i
	}
	req := Request{"demo", payload}
	v, err := json.Marshal(&req)
	if err != nil {
		log.Fatalln(err)
	}

	return string(v)
}

func processRequestLowPerf(reqs []string) []string {
	var reps []string
	for _, req := range reqs {
		reqObj := &Request{}
		json.Unmarshal([]byte(req), reqObj)
		ret := ""
		for _, e := range reqObj.Payload {
			ret += strconv.Itoa(e) + ","
		}
		repObj := &Response{reqObj.TransactionId, ret}
		repJson, err := json.Marshal(&repObj)
		if err != nil {
			log.Fatal(err)
		}
		reps = append(reps, string(repJson))
	}
	
	return reps
}

func processRequestHighPerf(reqs []string) []string {
	var reps []string
	for _, req := range reqs {
		reqObj := &Request{}
		reqObj.UnmarshalJSON([]byte(req))

		var buf strings.Builder
		for _, e := range reqObj.Payload {
			buf.WriteString(strconv.Itoa(e))
			buf.WriteString(",")
		}
		repObj := &Response{reqObj.TransactionId, buf.String()}
		repJson, err := repObj.MarshalJSON()
		if err != nil {
			log.Fatal(err)
		}
		reps = append(reps, string(repJson))
	}

	return reps
}