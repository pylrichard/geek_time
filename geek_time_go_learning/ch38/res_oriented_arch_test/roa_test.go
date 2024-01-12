package res_oriented_arch_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	hr "github.com/julienschmidt/httprouter"
)

type Employee struct {
	Id		string	`json:"id"`
	Name 	string	`json:"name"`
	Age 	int 	`json:"age"`
}

var employeeDb map[string]*Employee

func init() {
	employeeDb = map[string]*Employee{}
	employeeDb["Mike"] = &Employee{"e-1", "Mike", 35}
	employeeDb["Rose"] = &Employee{"e-2", "Rose", 40}
}

func Index(w http.ResponseWriter, r *http.Request, _ hr.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func GetEmployeeByName(w http.ResponseWriter, r *http.Request, ps hr.Params) {
	name := ps.ByName("name")
	var (
		ok			bool
		info		*Employee
		infoJson	[]byte
		err			error
	)

	if info, ok = employeeDb[name]; !ok {
		w.Write([]byte("{\"error\":\"Not Found\"}"))

		return
	}
	if infoJson, err = json.Marshal(info); err != nil {
		w.Write([]byte(fmt.Sprintf("{\"error\":,\"%s\"}", err)))

		return
	}

	w.Write(infoJson)
}

func TestROA(t *testing.T) {
	r := hr.New()
	r.GET("/", Index)
	r.GET("/employee/:name", GetEmployeeByName)

	t.Fatal(http.ListenAndServe(":8080", r))
}