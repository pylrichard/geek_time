package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeAndValue(t *testing.T) {
	var n int64 = 10
	t.Log(reflect.TypeOf(n), reflect.ValueOf(n))
	t.Log(reflect.ValueOf(n).Type())
}

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown", t)
	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(f)
	CheckType(&f)
}

type Customer struct {
	CookieId	string
	Name		string
	Age			int
}

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}
	//t.Log(a == b)
	t.Log("a == b?", reflect.DeepEqual(a, b))

	i1 := []int{1, 2, 3}
	i2 := []int{1, 2, 3}
	i3 := []int{2, 3, 1}

	t.Log("i1 == i2?", reflect.DeepEqual(i1, i2))
	t.Log("i1 == i3?", reflect.DeepEqual(i1, i3))

	c1 := Customer{"1", "Mike", 40}
	c2 := Customer{"1", "Mike", 40}
	t.Log(c1 == c2)
	t.Log(reflect.DeepEqual(c1, c2))
}

type Employee struct {
	EmployeeId	string
	Name		string `format:"normal"`
	Age			int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}
	t.Logf("Name: value(%[1]v), Type(%[1]T)", reflect.ValueOf(*e).FieldByName("Name"))
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'Name' field.")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}
	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("Updated ", e)
}