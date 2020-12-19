package main

import "fmt"

func main() {
	variadicDouble("um numero", 1)
	variadicDouble("dois", 4, 6)
	variadicDouble("cinco", 4, 6, 88, 12, 90)
}

// A variadic function is a function that accepts a variable number of argument values.
//  It is good to use a variadic function when the number of arguments of a specified type is unknown
func variadicDouble(name string, number ...int) {
	fmt.Println(name)
	fmt.Println(number)
}

// The preceding function is an example of what a variadic function looks like.
// The three dots (…) in front of the type is called a pack operator
