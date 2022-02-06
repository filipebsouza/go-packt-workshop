package main

import (
	"flag"
	"log"
	"sync"
	"sync/atomic"
)

// When multiple Goroutines race to use the same variable
// We could use atomic for safely modify variables across Goroutines
func atomicSum(from, to int, wg *sync.WaitGroup, res *int32) {
	for i := from; i <= to; i++ {
		atomic.AddInt32(res, int32(i))
	}

	wg.Done()
	return
}

func sum(from, to int, wg *sync.WaitGroup, res *int32) {
	for i := from; i <= to; i++ {
		*res = *res + int32(i)
	}

	wg.Done()
	return
}

var atomicFlag = -1

func main() {
	s1 := int32(0)
	wg := &sync.WaitGroup{}
	wg.Add(4)

	if atomicFlag == -1 {
		flag.IntVar(&atomicFlag, "atomic", -1, "Use atomic main or no")
		flag.Parse()
	}

	if atomicFlag == 1 {
		go atomicSum(1, 25, wg, &s1)
		go atomicSum(26, 50, wg, &s1)
		go atomicSum(51, 75, wg, &s1)
		go atomicSum(76, 100, wg, &s1)
	} else {
		go sum(1, 25, wg, &s1)
		go sum(26, 50, wg, &s1)
		go sum(51, 75, wg, &s1)
		go sum(76, 100, wg, &s1)
	}
	wg.Wait()

	log.Println(s1)
}
