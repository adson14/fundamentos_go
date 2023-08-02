package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo struct")
	var u usuario
	u.nome = "Adson"
	u.idade = 20
	fmt.Println(u)

	endereco1 := endereco{logradouro: "Rua ABC", numero: 10}

	//INferencia de tipos
	var u2 = usuario{
		nome:     "Adson",
		endereco: endereco1,
	}
	fmt.Println(u2)
}
