package main

import (
	"fmt"
)

//Speaker ...
//Duck typing is a test in computer programming: "If it looks like a duck, swims like a duck,
//and quacks like a duck, then it must be a duck." If a type matches an interface, then you can use
//that type wherever that interface is used.
type Speaker interface {
	Speak() string
}
type cat struct {
}

func main() {
	c := cat{}
	chatter(c)
}
func (c cat) Speak() string {
	return "Purr Meow"
}
func chatter(s Speaker) {
	fmt.Println(s.Speak())
}

// Anything that matches the Speak() method can be a Speaker{} interface.
// When implementing an interface, we are essentially conforming to that interface
// by having the required method sets.

// cat matches the Speak() method of the Speaker{} interface, so a cat is a Speaker{}
