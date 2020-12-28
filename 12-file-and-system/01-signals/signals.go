package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool)
	signal.Notify(sigs, syscall.SIGINT)
	go func() {
		for {
			s := <-sigs
			switch s {
			case syscall.SIGINT:
				fmt.Println()
				fmt.Println("My process has been interrupted. Someone might of pressed CTRL-C")
				fmt.Println("Some clean up is occuring")
				done <- true
			}
		}
	}()
	fmt.Println("Program is blocked until a signal is caught")
	<-done
	fmt.Println("Out of here")
}
