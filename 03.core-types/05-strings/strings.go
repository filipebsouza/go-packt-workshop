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
	safeLoopWithStrings()
}

func safeLoopWithStrings() {
	username := "Sir_King_Über"
	// Length of a string
	fmt.Println("Bytes:", len(username))
	fmt.Println("Runes:", len([]rune(username)))
	// Limit to 10 characters
	fmt.Println(string(username[:10]))
	fmt.Println(string([]rune(username)[:10]))
}
