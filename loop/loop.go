package main

import (
	"fmt"
	//"time"
)

func main() {
	i := 0

for i < 5 {
	fmt.Println("Imprimindo i. i = ", i)
	i++
	//time.Sleep(time.Second)
}
fmt.Println(i)

for j := 0; j < 5; j++ {
	fmt.Println("incrementando j")
}

for a := 0; a < 15; a+=5 {
	fmt.Println(a)
}

for i := 1; i <= 3; i++ {
	for j := 1; j <= 3; j++ {
		fmt.Printf("(%d, %d) ", i, j)
	}
	fmt.Println()  // Nova linha após cada ciclo do loop interno
}
nomes := []string{"Tais", "Tiago", "Diana"}

for indice, nome := range  nomes {
	fmt.Println(indice, nome)
}

for _, nome := range  nomes {
	fmt.Println(nome)
}

nums := []int{1,2,3}

for indice, nome := range nomes {
	for _, num := range nums {
		fmt.Println(indice, nome, num)
	}
}

//imprime seguindo ordem alfabetica
palavras := map[string] string {
	"aprimeira": "palavra 1",
	"bsegunda": "palavra 2",
	}

	for chave, valor := range palavras {
		fmt.Println(chave, valor)
	}

	mapaAninhado := map[string]map[string]string{
        "chave1": {
            "subchave1": "valor1",
            "subchave2": "valor2",
        },
        "chave2": {
            "subchave3": "valor3",
            "subchave4": "valor4",
        },
    }

    // Loop for aninhado para iterar sobre o mapa aninhado
    for chaveExterna, mapaInterno := range mapaAninhado {
        fmt.Printf("Chave Externa: %s\n", chaveExterna)

        for subchave, valor := range mapaInterno {
            fmt.Printf("  Subchave: %s, Valor: %s\n", subchave, valor)
        }
    }

	// for {
	// 	fmt.Println("Executando infinitamente")
	// 	time.Sleep(time.Second)
	// }
}

