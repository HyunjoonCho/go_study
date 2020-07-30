package main

import "fmt"

func main(){
	// var num int: OK when L9 num = ~~, NOT num := ~~ 
	// var num: error = unexpected newline, expecting type
	// => need its value to infer type
	if num := 12; num<0{ //cannot skip brace{} even for one-line code
		fmt.Println(num, "is negative")
	} else if(num<10){
		fmt.Println(num, "is one-digit number")
	} else{
		fmt.Println(num, "is multi-digits")
	}
}