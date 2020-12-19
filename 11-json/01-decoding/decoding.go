package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	func1()
	fmt.Println("====================")
	func2()
}

func func1() {
	type greeting struct {
		Message string
	}
	data := []byte(` 
	{ 
	  "message": "Greetings fellow gopher!" 
	} 
  	`)
	var v greeting
	err := json.Unmarshal(data, &v) //Decoding
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v.Message)
}

func func2() {
	type greeting struct {
		SomeMessage string `json="message"`
	}
	data := []byte(` 
	{ 
	  "message": "Greetings fellow gopher!" 
	} 
  	`)
	var v greeting
	err := json.Unmarshal(data, &v) //Decoding
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v.SomeMessage)
}

//Struct Tags
type person struct {
	LastName string `json:"lname"`
}
