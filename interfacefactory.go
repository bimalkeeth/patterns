package main

import "fmt"

type PersonFour interface {
	SayHello()
}

type personFour struct {
	name string
	age int
}

func(p *personFour)SayHello(){
	fmt.Printf("Hi,my name is %s,Iam %d years old\n",p.name,p.age)
}

func NewPersonFour(name string, age int) PersonFour {
	return &personFour{name: name, age: age}
}

func main() {
	p:=NewPersonFour("Bimal",50)
	p.SayHello()
}
