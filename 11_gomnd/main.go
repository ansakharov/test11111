package main

import "fmt"

// const `multiplier` is unused (unused)
const multiplier = 666

func main() {
	in := 100

	fmt.Println(MultipleInput(in))
}

func MultipleInput(in int) int {
	//return in * 666 // Magic number: 666, in <return> detected (gomnd)
	return in * multiplier
}
