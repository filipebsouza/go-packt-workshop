package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write([]byte("Using Write function.\n"))
	f.WriteString("Using Writestring function.\n")

	// Escrevendo passando as permissões
	message := []byte("Look!")
	err = ioutil.WriteFile("test.txt", message, 0644)
	if err != nil {
		fmt.Println(err)
	}

	criarVerificarSeArquivoExiste()
}

func criarVerificarSeArquivoExiste() {
	file, err := os.Stat("junk.txt") //verificar se arquivo existe
	if err != nil {
		if os.IsNotExist((err)) {
			fmt.Println("junk.txt: File does not exist!")
			fmt.Println(file)
		}
	}
	fmt.Println()
	file, err = os.Stat("test.txt")
	if err != nil {
		if os.IsNotExist((err)) {
			fmt.Println("test.txt: File does not exist!")
		}
	}
	fmt.Printf(
		"file name: %s\nIsDir: %t\nModTime: %v\nMode: %v\nSize: %d\n",
		file.Name(),
		file.IsDir(),
		file.ModTime(),
		file.Mode(),
		file.Size())
}
