package main

import (
	"bytes"
	"log"
	"os"
	"testing"
)

func Test_atomicMain(t *testing.T) {
	for i := 0; i < 1000; i++ {
		s := bytes.Buffer{}
		os.Args = append(os.Args, "--atomic=1")
		log.SetOutput(&s)
		log.SetFlags(0)
		main()

		if s.String() != "5050\n" {
			t.Errorf(s.String())
		}
	}
}

// Sometimes crash because of atomic
func Test_nonAtomicMain(t *testing.T) {
	for i := 0; i < 1000; i++ {
		s := bytes.Buffer{}
		os.Args = append(os.Args, "--atomic=0")
		log.SetOutput(&s)
		log.SetFlags(0)
		main()

		if s.String() != "5050\n" {
			t.Errorf(s.String())
		}
	}
}
