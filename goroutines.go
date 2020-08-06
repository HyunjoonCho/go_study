package main

import (
	"fmt"
	"time"
)

func f(msg string) {
	for i:=0; i<3; i++{
		fmt.Println(i, msg)
	}
}

func main(){
	f("direct")

	go func (msg string){
		fmt.Println(msg)
	} ("anonymous goroutine")

	go f("goroutine")	//output differs time to time

	time.Sleep(time.Second)
	fmt.Println("done")
}