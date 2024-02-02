package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64

}

func escreverArea(f forma) {
	fmt.Printf("a área da forma é %0.2f \n", f.area())
}

type retangulo struct {
	altura float64
	largura float64
}

type circulo struct {
	raio float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
	//pi * raio ao quadrado
}

func main() {
	retangulo := retangulo{11, 23}
	escreverArea(retangulo)

	circulo := circulo{34}
	escreverArea(circulo)
}