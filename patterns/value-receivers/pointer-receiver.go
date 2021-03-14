package main

import (
	"fmt"
)

// Foo not adding any pointer receiver to the struct makes it immutable as everytime we return a copy of the struct with values.
type Foo struct {
	name  string
	value int
}

func NewFoo(name string, value int) Foo {
	return Foo{
		name:  name,
		value: value,
	}
}

func (f Foo) GetName() string {
	return f.name
}

func (f Foo) SetName(name string) Foo {
	f.name = name
	return f
}

func main() {
	foo := NewFoo("name1", 1)
	fmt.Printf("%+v\n", foo)
	name := foo.GetName()
	fmt.Printf("name: %v\n", name)
	foo2 := foo.SetName("name2")
	fmt.Printf("%+v\n", foo2)
	name = foo2.GetName()
	fmt.Printf("name: %v\n", name)
}
