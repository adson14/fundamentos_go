package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendfo")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("tede@treste.com")
	fmt.Println(erro)
}
