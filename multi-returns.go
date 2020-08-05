package main

import "fmt"

func double()(int, int){
	return 3, 42
}

func main(){
	a, b := double()
	fmt.Println(a, b)
	_, c := double()
	fmt.Println(c)
}