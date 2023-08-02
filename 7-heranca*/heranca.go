package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	telefone  string
}

// Um tipo de "herança"
type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {

	// pessoa1 := pessoa {
	// 	nome: "Adson",
	// 	sobrenome: "Silva",
	// 	idade: 20,
	// 	telefone: "9999-9999",
	// }

	estudante1 := estudante{
		pessoa: pessoa{
			nome:      "Adson",
			sobrenome: "Silva",
			idade:     20,
			telefone:  "9999-9999",
		},
		curso:     "Engenharia de Software",
		faculdade: "IFSP",
	}

	fmt.Println(estudante1.nome)

}
