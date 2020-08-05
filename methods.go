package main 

import "fmt"

type rectangle struct{
	width int
	height int
}

func (r rectangle) area() int{
	return r.width*r.height
}

func (r *rectangle) perimeter() int{
	return 2*r.width + 2*r.height
}

func main(){
	r := rectangle{3, 4}
	fmt.Println("area:",r.area())
	fmt.Println("peri:",r.perimeter())

	rp := &rectangle{5,12}
	fmt.Println("area:",rp.area())
	fmt.Println("peri:",rp.perimeter())
}