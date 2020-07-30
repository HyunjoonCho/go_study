package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)
	a = "now?"
	fmt.Println(a)
	
	var b, c int = 1, 2
	fmt.Print(b, c) // just Print = no line shift
	d := "apple"
	fmt.Println(d)
}
