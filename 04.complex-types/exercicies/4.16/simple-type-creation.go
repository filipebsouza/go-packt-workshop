package main

import "fmt"

func main() {
	id1, id2, id3 := getIds()

	fmt.Println(id1)
	fmt.Println(id2)
	fmt.Println(id3)
}

type id int

func getIds() (id, id, id) {
	return 1, 2, 3
}
