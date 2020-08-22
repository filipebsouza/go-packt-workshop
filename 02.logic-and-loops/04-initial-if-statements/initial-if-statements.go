package main

import (
	"errors"
	"fmt"
)

func main() {
	input := 21

	if err := validate(input); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(input, "is odd")
	}
}

func validate(input int) error {
	if input < 0 {
		return errors.New("input can't be a negative number")
	} else if input > 100 {
		return errors.New("input can't be over 100")
	} else if input%7 == 0 {
		return errors.New("input can't be divisible by 7")
	} else {
		return nil
	}
}
