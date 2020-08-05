package main 

import (
	"fmt"
	"math"
)

type geometry interface{
	area() float64
	peri() float64
}

//no explicit connection between struct & interface
//rather, contain all members of interface as method
//https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go 
type rect struct{	
	width, height float64
}

type circle struct{
	radius float64
}

func (r rect) area() float64{
	return r.height*r.width
}
func (r rect) peri() float64{
	return 2*(r.height+r.width)
}

func (c circle) area() float64{
	return math.Pi*c.radius*c.radius
}
func (c circle) peri() float64{
	return 2*math.Pi*c.radius
}

func measure(g geometry){
	fmt.Println(g)
	fmt.Println("area:",g.area())
	fmt.Println("peri:",g.peri())
}

func main(){
	r := rect{3, 4}
	c := circle{5}

	measure(r)
	measure(c)
}