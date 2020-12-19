package main

import "fmt"

//Pessoa ...
type Pessoa interface {
	Fala() string
	Come(comidas []string)
}

type filipe struct {
}

func (f filipe) Fala() string {
	return "Opa"
}

func (f filipe) Come(comidas []string) {
	for _, item := range comidas {
		fmt.Printf("Comida: %v", item)
	}
}

func main() {
	f := filipe{}
	fmt.Println(f.Fala())
	f.Come([]string{"Açai", "Banana", "Bife"})
}
