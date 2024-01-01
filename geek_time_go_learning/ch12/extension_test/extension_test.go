package extension_test_test

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Println("pet speak")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Printf("pet speak to %v\n", host)
}

type Dog1 struct {
	p *Pet
}

//Go没有继承，复合需要自行实现相应方法
func (d1 *Dog1) Speak() {
	//d1.p.Speak()
	fmt.Println("dog1 speak")
}

func (d1 Dog1) SpeakTo(host string) {
	//d1.p.SpeakTo(host)
	d1.Speak()
	fmt.Printf("dog1 speak to %v\n", host)
}

func TestDog1(t *testing.T) {
	dog1 := new(Dog1)
	dog1.SpeakTo("dog1")
}

type Dog2 struct {
	//匿名类型嵌入不需要实现相应方法
	Pet
}

//不支持重载和LSP
func (d2 *Dog2) Speak() {
	fmt.Println("dog2 speak")
}

func TestDog2(t *testing.T) {
	dog2 := new(Dog2)
	dog2.SpeakTo("dog2")

	//不支持LSP
	//var dog3 *Dog2 = new(Dog2)
	//var p *Pet = (*Pet)(dog3)
}