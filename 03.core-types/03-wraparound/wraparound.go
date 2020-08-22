package main

import (
	"fmt"
	"math"
)

func main() {
	var a int8 = 125
	var b uint8 = 253

	for i := 0; i < 5; i++ {
		a++
		b++
		fmt.Println(i, ")", "int8", a, "uint8", b)
		// Quando passa do valor limite, ele volta para o menor valor e continua a sequência.
	}

	// get max value
	int64MaxValue := math.MaxInt64

	fmt.Println(int64MaxValue)

	// The byte type in Go is just an alias for uint8
}
