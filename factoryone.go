package main


type PersonThree struct {
	Nme string
	Age int
	EyeCount int
}

func NewPersonThree(name string,age int) *PersonThree{
	if age < 16 {

	}
	return &PersonThree{Nme: name,Age: age}
}


func main() {
	
}
