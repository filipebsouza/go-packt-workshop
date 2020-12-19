package main

import (
	"log"
	"os"
)

//The Go log package has a function called SetFlags that allows us to be more specific.
func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	name := "Thanos"
	log.Println("Demo app")
	log.Printf("var: %#+v\n", name)
	log.Print("Run")
	os.Exit(0)
	//log.Fatal()
}

// Using the log package, we can also log fatal errors.
// The Fatal(), Fatalf(), and Fatalln() functions are similar to Print(), Printf(), and Println().
// The difference is after the log Fatal() functions are followed by an os.Exit(1) a system call.
// The log package also has the following functions: Panic, Panicf, and Panicln.
// The difference between the Panic() functions and the Fatal function is that the Panic
// functions are recoverable. When using the Panic functions, you can use the defer() function,
// whereas when using the Fatal functions, you cannot.
// As stated earlier, the Fatal functions call os.Exit();
// a defer function will not be called when an os.Exit() gets called.
