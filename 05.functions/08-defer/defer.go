// hen using defer  statements in front of functions, the execution follows the order of First In Last Out  ( FILO ).
package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Execução de função deferiada para o fim do método, depois do retorno")
	}()

	fmt.Println("Fluxo normal")
}
