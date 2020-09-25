package main

import "fmt"

func main() {
	comp1, comp2, comp3 := compArrays()
	fmt.Println("[5]int == [5]int{0}       :", comp1)
	fmt.Println("[5]int == [...]int{0, 0, 0, 0, 0}:", comp2)
	fmt.Println("[5]int == [5]int{0, 0, 0, 0, 9} :", comp3)
}

func compArrays() (bool, bool, bool) {
	var arr1 [5]int
	arr2 := [5]int{0}
	arr3 := [...]int{0, 0, 0, 0, 0}
	arr4 := [5]int{0, 0, 0, 0, 9}

	// Arrays usando chave/valor:
	_ = [...]int{9: 0} // _ Blank Identifier
	_ = [10]int{1, 9: 10, 4: 5}

	// Note: Using len in a loop- In other languages,
	// it's not efficient to count the number of elements on each iteration of a loop.
	// In Go, it's okay. The Go runtime tracks the length of the array internally,
	// so it doesn't count the items when you call len.
	// This feature is also true for the other collection types, that is, slice and map.

	var arr8 []int
	arr8 = append(arr8, 7, 8)

	return arr1 == arr2, arr1 == arr3, arr1 == arr4
}
