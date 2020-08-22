package main

import "fmt"

//Go has a single type to represent some text, string.
func main() {
	//Raw – defined by wrapping text in a pair of `
	//Interpreted – defined by surrounding the text in a pair of "

	comment1 := `This is the BEST 
				thing ever!`
	comment2 := `This is the BEST\nthing ever!`
	comment3 := "This is the BEST\nthing ever!"
	fmt.Print(comment1, "\n\n")
	fmt.Print(comment2, "\n\n")
	fmt.Print(comment3, "\n")
}
