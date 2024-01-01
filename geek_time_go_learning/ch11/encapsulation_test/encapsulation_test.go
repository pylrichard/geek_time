package encapsulation_test_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id		string
	Name	string
	Age		int
}

//*Employee避免了实例成员的值复制
func (e *Employee) String() string {
	fmt.Printf("addr: %x\n", unsafe.Pointer(&e.Name))

	return fmt.Sprintf("id:%s-name:%s-age:%d", e.Id, e.Name, e.Age)
}

func TestEmployee(t *testing.T) {
	//e1 := Employee{"0", "Bob", 20}
	e2 := Employee{Name: "Mike", Age: 30}
	//new()返回实例指针&Employee{}
	e3 := new(Employee)
	e3.Id = "2"
	e3.Name = "Rose"
	e3.Age = 22

	//t.Log(e1)
	//t.Logf("e1 is %T", e1)
	fmt.Printf("addr: %x\n", unsafe.Pointer(&e2.Name))
	t.Log(e2)
	t.Log(e2.String())
	//t.Log(e3)
	//t.Logf("e3 is %T", e3)
}

type Rectangle struct {
	width	float64
	height	float64
}

//值接收者
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

//指针接收者
func (r *Rectangle) Scale(factor float64)  {
	r.width *= factor
	r.height *= factor
}

func TestRectangle(t *testing.T) {
	rect := Rectangle{width: 10, height: 15}
	t.Log(rect.Area())
	rect.Scale(2)
	t.Log(rect.Area())
}