package main

import(
	"fmt"
	"time"
)

func main(){
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("Uh-oh it's weekday")
	}

	t := time.Now()
	switch{
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	
	whatAmI := func(i interface{}){	// first-class function i.e. function as a variable
		// REF: https://golangbot.com/first-class-functions/
		switch t := i.(type) {
		case bool: 
			fmt.Println("I am bool")
		case int:
			fmt.Println("I am int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

// func whatAmI(i interface{}){
// 	switch t := i.(type) {
// 	case bool: 
// 		fmt.Println("I am bool")
// 	case int:
// 		fmt.Println("I am int")
// 	default:
// 		fmt.Printf("Don't know type %T\n", t)
// 	}
// }