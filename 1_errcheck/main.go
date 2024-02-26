package main

import (
	"fmt"
	"github.com/pkg/errors"
)

const secondDegree = 2

func someFunc(in int) (int, error) {
	if in%2 == 0 {
		return 0, errors.New("some_err")
	}

	return in ^ secondDegree, nil
}

//nolint:gomnd
func main() {
	//res, _ := someFunc(2)
	//fmt.Println(res)
	// Check error
	res, err := someFunc(2)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(res)
}
