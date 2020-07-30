package main

import "fmt"

func main(){
	i := 1
	for i<7 { 
		fmt.Println(i)
		i = i + 1
	}
	for j:=7; j<9;j++{
		fmt.Println(j)
	}
	for{
		fmt.Println("loop")
		break
	}
	for n:=0; n<5; n++{
		if(n%2==1){
			continue
		}
		fmt.Println(n)
	}
}