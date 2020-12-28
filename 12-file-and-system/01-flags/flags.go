package main

import (
	"flag"
	"fmt"
)

// Flags podem ser usadas para passar parâmetros em programas de linha de comando
func main() {
	i := flag.Int("age", -1, "your age")
	n := flag.String("name", "", "your first name")
	b := flag.Bool("married", false, "are you married?")
	flag.Parse()
	fmt.Println("Name: ", *n)
	fmt.Println("Age: ", *i)
	fmt.Println("Married: ", *b)
}

// Usando flags
//go run . -name=John
//go run . -name=John –age 42 -married true
//go run . -h
