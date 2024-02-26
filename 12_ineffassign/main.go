package main

import "fmt"

func main() {
	// z := 100
	a := 0
	b := 0
	// c := 0

	fmt.Println(a)
	fmt.Println(b)

	//a = 1 // This reassignment is wasted, because never used afterwards. Wastedassign find this
	//b = 1 // This reassignment is wasted, because reassigned without use this value. Wastedassign find this
	b = 2

	fmt.Println(b)
}
