package http_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func httpHandler() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	http.HandleFunc("/time/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\": \"%s\"}", t)
		w.Write([]byte(timeStr))
	})

	http.ListenAndServe(":8080", nil)
}

func TestHttp(t *testing.T) {
	httpHandler()
}