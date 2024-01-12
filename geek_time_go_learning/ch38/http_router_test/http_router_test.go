package http_router_test

import (
	"fmt"
	hr "github.com/julienschmidt/httprouter"
	"net/http"
	"testing"
)

func Index(w http.ResponseWriter, r *http.Request, _ hr.Params) {
	fmt.Fprintf(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps hr.Params) {
	fmt.Fprintf(w, "Hello %s!\n", ps.ByName("name"))
}

func TestHttpRouter(t *testing.T) {
	r := hr.New()
	r.GET("/", Index)
	r.GET("/hello/:name", Hello)

	t.Fatal(http.ListenAndServe(":8080", r))
}