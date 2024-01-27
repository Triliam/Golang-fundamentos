package main

import "fmt"

func main() {
// [chave]valor
	usuario := map[string]string {
		"nome": "Diana",
		"sobrenome": "Balque",
		"email": "db@gmail.com",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]int {
		"idade": 19,
		"numero": 22,
	}
	fmt.Println(usuario2)

	//Maps aninhados
	usuario3 := map[string]map[string]string {
		"nome": {
			"primeiro": "Tais",
			"segundo": "Bueno",
			"ultimo": "Vidotto",
		},
		"curso": {
			"semestre": "primeiro",
		},
	}
		fmt.Println(usuario3)

		// deletar um map de dentro dos aninhados
		delete(usuario3, "curso")
		fmt.Println(usuario3)

		// add um map dentro do outro
		usuario3["signo"] = map[string]string {
			"nomeSigno": "Virgem",
		}

		fmt.Println(usuario3)
	}
	// Estrutura com chave e valor, chave do mesmo tipo e imutavel, valor do mesmo tipo e imutavel, tipo da chave pode ser diferente do tipo do valor
