package main

import "fmt"

// START OMIT

type SomeInterface interface {
	someFunc()
}
type Impl struct{}

func (_ Impl) someFunc() { fmt.Println("real") }

type FakeImpl struct{}

func (_ FakeImpl) someFunc() { fmt.Println("fake") }

func doStuff(si SomeInterface) {
	si.someFunc()
}

func main() {
	realImpl := &Impl{}
	fakeImpl := &FakeImpl{}
	doStuff(realImpl)
	doStuff(fakeImpl)
}

// END OMIT
