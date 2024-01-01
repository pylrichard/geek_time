package interface_test

import "testing"

type Programmer interface {
	Write() string
}

type GoProgrammer struct {}

func (gp *GoProgrammer) Write() string {
	return "Hello GoProgrammer"
}

func TestInterface(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.Write())
}