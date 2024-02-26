package main

import (
	"errors"
	"fmt"
)

//nolint:errname
var MyError = errors.New("my error")

func getMyError() error {
	return MyError
}

func errorWrapper(err error) error {
	return fmt.Errorf("wrapped error: %w", err)
}

func main() {
	err := getMyError()
	//if err == MyError {
	//	fmt.Println("err is MyError")
	//} else {
	//	fmt.Println("err NOT MyError")
	//}

	if errors.Is(err, MyError) {
		fmt.Println("err is MyError")
	} else {
		fmt.Println("err NOT MyError")
	}

	wrappedErr := errorWrapper(err)
	//if wrappedErr == MyError {
	//	fmt.Println("wrappedErr is MyError")
	//} else {
	//	fmt.Println("wrappedErr NOT MyError")
	//}

	if errors.Is(wrappedErr, MyError) {
		fmt.Println("wrappedErr is MyError")
	} else {
		fmt.Println("wrappedErr NOT MyError")
	}
}
