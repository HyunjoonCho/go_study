package main 

import "fmt"

func main(){
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func(){
		for{
			j, more := <-jobs
			if more{
				fmt.Println("received job", j)
			} else{
				fmt.Println("jobs received")
				done <- true
				return
			}
		}
	}()

	for j:=0; j<3; j++{
		jobs <- j
		fmt.Println("sent job", j)	
		//send & print NOT atomic -> timing may differ
	}
	close(jobs)
	fmt.Println("jobs sent")
	<-done	//this channel receiver waits for the function call is finished
}