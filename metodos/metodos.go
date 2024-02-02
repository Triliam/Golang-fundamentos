package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
}

//metodos associados a struct

func (u usuario) salvar() {
	fmt.Printf("metodo associado a struct usuario que salva os dados dos usuarios. Salvando o nome %s \n ", u.nome)
}

func (u usuario) maiorIdade() bool {
	return u.idade >= 18

	// if u.idade >= 18 {
	// 	return true
	// } else {
	// 	return false
	// }
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario {"Ursula", 20}
	usuario1.salvar()

	usuario2 := usuario{"Pedro", 33}
	usuario2.salvar()

	usuario3 := usuario{"Gabi", 17}
	usuario3.salvar()

	idade1 := usuario1.maiorIdade()
	fmt.Println(idade1)

	idade2 := usuario2.maiorIdade()
	fmt.Println(idade2)

	idade3 := usuario3.maiorIdade()
	fmt.Println(idade3)
}