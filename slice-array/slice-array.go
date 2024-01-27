package main

import "fmt"

func main() {

	//arrays sao do mesmo tipo e de tamanho definido/fixo
	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [3]string{"Posi0", "Posi1", "Posi2"}
	fmt.Println(array2)

	// capacidade/tamanho de acordo com o que for definido na atribuição (qdt de item passados nas posições)
	array3 := [...]int{1, 2, 3}
	fmt.Println(array3)

	//slice é mais flexivel em questao de capacidade/tamanho
	slice := []int{11, 12, 13, 14, 15}
	fmt.Println(slice)
	slice = append(slice, 16)
	fmt.Println(slice)

	//Um slice é um pedaço de array
	slice2 := array2[0:2]
	fmt.Println(slice2)

	array2[1] = "Alterei a posicao"
	fmt.Println(slice2)

	//ARRAYS INTERNOS

	//Função make: aloca espaço na memoria 
	//3 parametros: tipo, tamanho, capacidade
	//cria um array de 15 posições e pega as primeiras 10 posições e mostra pelo slice
	//make cria um array interno (qnd nao se cria um slice relacionado a algum array declarado)
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade
	slice3 = append(slice3, 4)
	slice3 = append(slice3, 3)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) 
	fmt.Println(cap(slice3)) 

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	slice4 = append(slice4, 10)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}

// geralmente usa mais o slice pela flexibilidade de capacidade e tamanho. 