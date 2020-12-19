package main

import "fmt"

func main() {
	typeCheck(1)
	typeCheck("opa")
	typeCheck(true)
	typeCheck(23.5)

	fmt.Println("---------------")
	func() {
		var inteiro interface{} = 12
		v, isValid := inteiro.(string)
		fmt.Println(v, isValid)
	}()
}

func typeCheck(i ...interface{}) {
	for _, x := range i {
		switch v := x.(type) {
		case int:
			fmt.Printf("%v is int\n", v)
		case string:
			fmt.Printf("%v is a string\n", v)
		case bool:
			fmt.Printf("a bool %v\n", v)
		default:
			fmt.Printf("Unknown type %T\n", v)
		}
	}
}
