package main

import (
	"fmt"
	"time"
)

func escrever(texto string) {
	for i := 0; i <= 5; i++{
		fmt.Println(texto)
	}
}

func main() {
	go escrever("primeiro texto")
	escrever("segundo texto")
	time.Sleep(time.Second)
}