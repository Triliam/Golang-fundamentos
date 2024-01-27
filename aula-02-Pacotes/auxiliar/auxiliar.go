package auxiliar

import (
	"fmt"	
)

// Escrever registra uma mensagem e tem visibilidade publica (Inicial maiuscula), e chama a função escrever2
func Escrever() {
	fmt.Println("Escrevendo uma mensagem.")
	escrever2()
}