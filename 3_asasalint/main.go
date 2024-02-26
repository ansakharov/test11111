package main

import "fmt"

func innerFunc(args ...any) int {
	fmt.Printf("innerFunc: %v\n", args)

	return len(args)
}

func outerFunc(args ...any) int {
	fmt.Printf("outerFunc: %v\n", args)

	// pass []any as any to func innerFunc func(args ...any) int (asasalint)
	return innerFunc(args...)
}

func main() {
	// 1
	// nolint:gomnd
	fmt.Println(outerFunc(1, 2, 3, 4))
}
