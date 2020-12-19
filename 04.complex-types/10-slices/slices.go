package main

import (
	"fmt"
	"reflect"
)

func main() {
	// make(<sliceType>, <length>, <capacity>)
	s1 := make([]int, 5, 10)

	fmt.Println(reflect.TypeOf(s1)) // Tipo da variavel

	s1 = []int{45, 32, 888, 889, 97}
	s2 := append(s1, []int{6: 11, 12, 13}...)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1)

	//Slice de um slice
	numeros1 := []int{1, 2, 3, 4}
	fmt.Println(numeros1[0:3]) //Resultado: [1 2 3]
	//numeros1[0:3] é assim que fatiamos a nossa slice, vamos indicar onde nossa fatia vai iniciar e onde ela irá terminar.
	//Um detalhe, indicamos que ela termina no índice 3 (número 4), mas apenas mostramos até o índice 2 (número 3).
	//Isso é padrão, temos que ir um índice a mais do que queremos mostrar.
	//Pode ser um pouco confuso mesmo.

	///O nosso slice vai do índice 0 ao 3, mas no fmt.Println(numeros1[1:4]) indicamos o índice 4.
	//Queremos mostrar o último elemento da slice (índice 3), mas por padrão devemos adicionar mais um,
	//então temos que indicar o índice 4.
	//O fmt.Println(numeros1[0: ]) também mostra o último elemento, mas não indicamos nenhum índice.
	// Essa é outra opção, deixamos o campo vazio.

	// Can you see something new here?
	// This uses the seldom-used slice range notation of <slice>[<low>:<high>:<capacity>].
	s2 = append(s1[:4:5], s1...)

	fmt.Println(s2)
}
