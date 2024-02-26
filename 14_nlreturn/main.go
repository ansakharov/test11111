package main

import "fmt"

const input = 100
const divider = 2

func main() {
	result := halfEvenNumber(input)

	fmt.Println(result)
}

func halfEvenNumber(in int) int {
	if in%2 == 0 {
		_ = in // return with no blank line before (nlreturn)

		return in / divider
	}

	return in
}
