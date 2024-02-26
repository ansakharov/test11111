package main

import (
	"errors"
	"fmt"
)

// the variable name `ErrCustom` should conform to the `ErrXxx` format (errname)
var ErrCustom = errors.New("my custom error")

func main() {
	err := ErrCustom
	if err != nil {
		fmt.Println(err)
	}
}
