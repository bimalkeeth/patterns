package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAddress,City,Country string
}

type PersonFive struct{
	Name string
	Address *Address
	Friends []string
}
func(a *Address)DeepCopy() *Address{
	return &Address{
		City: a.City,
		Country: a.Country,
		StreetAddress: a.StreetAddress,
	}
}
func(p *PersonFive)DeepCopy() *PersonFive{
 q:=*p
 q.Address=p.Address.DeepCopy()
 copy(q.Friends,p.Friends)
 return &q
}

func(p *PersonFive)DeepCopyNext()*PersonFive{
	b:=bytes.Buffer{}
	e:=gob.NewEncoder(&b)
	_=e.Encode(p)
	d:=gob.NewDecoder(&b)
	result:=PersonFive{}
	_=d.Decode(&result)
	return &result
}


func main() {
	john :=PersonFive{"John",
		&Address{"123 London Road","London","UK"},[]string{"Chris","Matt"}}

    jane:=john.DeepCopy()
    jane.Name="Jane"
    jane.Address.StreetAddress="321 Baker Street"
    jane.Friends=append(jane.Friends,"Anjela")

    fmt.Println(jane,jane.Address)

    newcop:= jane.DeepCopyNext()
	fmt.Println(newcop,newcop.Address)
}
