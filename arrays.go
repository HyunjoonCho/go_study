package main

import "fmt"

func main(){
	var a [5]int
	fmt.Println("emp:", a)

	// a[5] = 100: invalid array index 5 (out of bounds for 5-element array)
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))
	b := [6]int{1,2,3,4,5} // fills zeros at the end
	fmt.Println("dcl:", b)

	// m := 2
	// n := 3
	// var twoD[m][n]int: non-constant array bound n -> slice?
	var twoD[2][3]int
	fmt.Printf("%T\n", twoD) //type = [2][3]int
	for i:=0; i<2; i++{
		for j:=0; j<3; j++{
			twoD[i][j] = i+j
		}
	}
	fmt.Println("2D:", twoD)
}