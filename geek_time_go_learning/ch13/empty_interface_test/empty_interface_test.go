package empty_interface_test_test

import (
	"fmt"
	"testing"
)

func GetInterfaceType(i interface{}) {
	if e, ok := i.(bool); ok {
		fmt.Println("Bool", e)
		return
	}

	switch v := i.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknown Type")
	}
}

func TestEmptyInterface(t *testing.T) {
	GetInterfaceType(true)
	GetInterfaceType(10)
	GetInterfaceType("10")
}