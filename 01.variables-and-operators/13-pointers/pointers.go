package main

import (
	"fmt"
	"time"
)

func main() {
	// Pointer
	var count1 *int
	// Variabvle with new
	count2 := new(int)
	countTemp := 5
	// Pointer with &
	count3 := &countTemp

	t := &time.Time{}

	if count1 != nil {
		fmt.Printf("count1: %#v\n", *count1)
	}
	if count2 != nil {
		fmt.Printf("count2: %#v\n", *count2)
	}
	if count3 != nil {
		fmt.Printf("count3: %#v\n", *count3) // Equivalente ao .Value
	}
	if t != nil {
		fmt.Printf("time : %#v\n", t.String())
	}
}
