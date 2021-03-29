package main

type InterfaceA interface {
	MethodA() string
	MethodB() string
}

type InterfaceB interface {
	GetA() InterfaceA
}

type InterfaceImplA struct{}

func (i *InterfaceImplA) MethodA() string {
	return "interfaceA-methodA"
}

func (i *InterfaceImplA) MethodB() string {
	return "interfaceA-methodB"
}

type InterfaceImplB struct{}

// This does not work, because the method signature should be exactly the same as defined in interface.
//func (i *InterfaceImplB) GetA() *InterfaceImplA {
//	return &InterfaceImplA{}
//}

func (i *InterfaceImplB) GetA() InterfaceA {
	return &InterfaceImplA{}
}

var _ InterfaceA = &InterfaceImplA{}
var _ InterfaceB = &InterfaceImplB{}
