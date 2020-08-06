package main 

import "fmt"

func ping(pings chan<- string, msg string){
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string){
	msg := <-pings
	pongs <- msg
}

func main(){
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	// DEADLOCK w/out buffer! WHY?
	// unbuffered = blocks till received / buffered = blocks till copied to buffer
	// https://stackoverflow.com/questions/18660533/why-does-the-use-of-an-unbuffered-channel-in-the-same-goroutine-result-in-a-dead
	ping(pings, "MESSAGE")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}