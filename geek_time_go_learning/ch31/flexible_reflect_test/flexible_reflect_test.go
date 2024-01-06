package flexible_reflect_test

import (
	"errors"
	"reflect"
	"testing"
)

type Employee struct {
	EmployeeId string
	Name       string `format:"normal"`
	Age        int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

type Customer struct {
	CookieId string
	Name     string
	Age      int
}

func fillBySettings(obj interface{}, settings map[string]interface{}) error {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return errors.New("the first param should be a pointer to the struct type")
	}
	//Elem()获取指针指向的值
	if reflect.TypeOf(obj).Elem().Kind() != reflect.Struct {
		return errors.New("the first param should be a pointer to the struct type")
	}
	if settings == nil {
		return errors.New("settings is nil")
	}

	var (
		field	reflect.StructField
		ok		bool
	)
	for k, v := range settings {
		if field, ok = (reflect.ValueOf(obj)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			valueOfObj := reflect.ValueOf(obj)
			valueOfObj = valueOfObj.Elem()
			valueOfObj.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}

	return nil
}

func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "Mike", "Age": 30}
	e := Employee{}
	if err := fillBySettings(&e, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(e)
	c := new(Customer)
	if err := fillBySettings(c, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)
}