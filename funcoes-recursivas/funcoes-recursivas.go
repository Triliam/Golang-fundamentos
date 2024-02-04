package main

import "fmt"

func fibonacci(num uint) uint {
	if num <= 1 {
		return num
	}

	return fibonacci(num-2) + fibonacci(num-1)
}

func main() {
	fmt.Println("Funções Recursivas")

	// 1 1 2 3 5 8 13

	posicao := uint(12)

	for i := uint(1); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}