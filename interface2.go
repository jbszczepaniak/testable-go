package main

import "fmt"

// START OMIT
type InterfaceToImplement interface {
	InterfaceA
	InterfaceB
	Run()
}

type InterfaceA interface {
	A()
}

type InterfaceB interface {
	B()
}

type SomeStruct struct {
	InterfaceA
	InterfaceB
}

// END OMIT

// START2 OMIT
func (ss SomeStruct) Run() {
	ss.A()
	ss.B()
}

type StructA struct{}

func (_ StructA) A() { fmt.Println("I am A") }

type StructB struct{}

func (_ StructB) B() { fmt.Println("I am B") }

// END2 OMIT

// START3 OMIT
type FakeA struct{}

func (_ FakeA) A() { fmt.Println("I am faked A") }

func main() {
	var legit InterfaceToImplement
	legit = &SomeStruct{InterfaceA: &StructA{}, InterfaceB: &StructB{}}
	legit.Run()

	var fake InterfaceToImplement
	fake = &SomeStruct{InterfaceA: &FakeA{}, InterfaceB: &StructB{}}
	fake.Run()
}

// END3 OMIT
