package main

import "fmt"

type person struct{
	name string
	age int
}

func newPerson(name string) *person{
	p:= person{name: name}
	p.age = 42
	return &p
}

func main(){
	fmt.Println(person{"Bob", 20})		//put fields in order
	fmt.Println(person{name:"Alice", age:35})
	fmt.Println(person{name:"Alfred"})	//must put "name"
	fmt.Println(newPerson("Joe"))

	s := person{"Sean", 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)
	fmt.Printf("%T %T\n", s, sp)

	sp.age = 51
	fmt.Println(sp.age)
}