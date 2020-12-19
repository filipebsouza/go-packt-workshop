package main

import (
	"errors"
	"fmt"
)

func main() {
	res, _ := doubler(5)
	fmt.Println("5 :", res)
	res, _ = doubler("yum")
	fmt.Println("yum :", res)
	_, err := doubler(true)
	fmt.Println("true:", err)
}

// Type Assertion
func doubler(v interface{}) (string, error) {
	if i, ok := v.(int); ok {
		return fmt.Sprint(i * 2), nil
	}

	if s, ok := v.(string); ok {
		return s + s, nil
	}

	return "", errors.New("Tipo não suportado")
}
