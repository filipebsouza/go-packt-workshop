package main

import "fmt"

func init() {
	fmt.Println("Inicializa o pacote, executa antes do main")
}

func main() {
	fmt.Println("Função main")
}
