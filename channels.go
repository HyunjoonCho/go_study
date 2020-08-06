package main 

import "fmt"

func main(){
	messages := make(chan string)

	go func(){ messages <- "channel works" }()

	msg := <-messages
	fmt.Println(msg)
}