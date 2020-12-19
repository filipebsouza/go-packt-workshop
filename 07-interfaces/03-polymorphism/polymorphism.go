//Polymorphism is the ability to appear in various forms

package main

import (
	"fmt"
)

//Speaker ...
type Speaker interface {
	Speak() string
}
type cat struct {
}
type dog struct {
}
type person struct {
	name string
}

func main() {
	c := cat{}
	d := dog{}
	p := person{name: "Heather"}
	saySomething(c, d, p)
}

func saySomething(say ...Speaker) {
	for _, s := range say {
		fmt.Println(s.Speak())
	}
}

func (c cat) Speak() string {
	return "Purr Meow"
}

func (d dog) Speak() string {
	return "Woof Woof"
}

func (p person) Speak() string {
	return "Hi my name is " + p.name + "."
}
