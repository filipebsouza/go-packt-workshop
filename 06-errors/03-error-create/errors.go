package main

import (
	"errors"
	"fmt"
)

var (
	errNumberNotBeNegative = errors.New("number not be negative")
)

func main() {
	result, err := sum(2, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	result, err = sum(-2, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func sum(n1, n2 int) (int, error) {
	if n1 < 0 || n2 < 0 {
		return 0, errNumberNotBeNegative
	}

	return n1 + n2, nil
}
