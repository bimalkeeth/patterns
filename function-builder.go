package main

import "fmt"

type PersonTwo struct {
	name,position string
}

type personMethod func(two *PersonTwo)
type PersonTwoBuilder struct {
	actions []personMethod
}

func(b *PersonTwoBuilder) Called(name string ) *PersonTwoBuilder{
	b.actions=append(b.actions, func(p *PersonTwo) {
		p.name=name
	})
	return b
}
func(b *PersonTwoBuilder)Build()*PersonTwo{
	p:=PersonTwo{}
	for _,a:=range b.actions{
		a(&p)
	}
	return &p
}

func main() {
	b:=PersonTwoBuilder{}
	p:=b.Called("Bimal Kaluarachchi").Build()
	fmt.Println(p.name)
}
