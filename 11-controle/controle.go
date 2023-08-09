package main

import "fmt"

func main() {
	fmt.Println("Estrutura de contorle")

	numero := 15

	if numero > 15 {
		fmt.Println("Maior")
	} else {
		fmt.Println("Menor")
	}

	// A variável é limitada ao escopo
	if outro := numero; outro > 15 {
		fmt.Println("Maior")
	} else if numero < -15 {
		fmt.Println("Menor")
	} else {
		fmt.Println("Igual")
	}

}
