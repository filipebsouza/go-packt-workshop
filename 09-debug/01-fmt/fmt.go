package main

import (
	"fmt"
	"log"
)

//The fmt.Printf() formats the string according to the verb and prints it to stdout.
// The function uses something called format verbs or sometimes called a format specifier.

func main() {
	teste := 1

	// Print do tipo da variavel
	fmt.Printf("teste is of type %T\n", teste)

	// Print do valor da variavel
	fmt.Printf("teste value %#v\n", teste)

	// Performar o erro
	log.Printf("teste value %#v\n", teste)
}
