package main 

import (
	"fmt"
	// "time"
)

func f(messages chan int, n int){
	messages<-n
}

func main(){
	messages := make(chan int, 5)

	for i:=0; i<5; i++{
		go f(messages, i)
	}

	for i:=0; i<5; i++{
		fmt.Println(<-messages)
		// go fmt.Println(~) does NOT PRINT the messages -> since not waited
	}

	// time.Sleep(time.Second)
}