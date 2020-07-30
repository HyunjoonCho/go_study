package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main(){
	fmt.Println(s)

	const n = 500000
	const d = 3.0/n
	fmt.Println(d)
	// fmt.Println(int64(d)): constant 6e-06 truncated to integer
	fmt.Println(math.Sin(n))
}