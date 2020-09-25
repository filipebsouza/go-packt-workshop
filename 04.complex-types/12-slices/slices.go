package main

import "fmt"

func main() {
	// make(<sliceType>, <length>, <capacity>)
	s1 := make([]int, 5, 10)
	s1 = []int{45, 32, 888, 889, 97}
	s2 := append(s1, []int{10: 11, 12, 13}...)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1)

	// Can you see something new here?
	// This uses the seldom-used slice range notation of <slice>[<low>:<high>:<capacity>].
	s2 = append(s1[:4:5], s1...)

	fmt.Println(s2)
}
