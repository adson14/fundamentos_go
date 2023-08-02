package main

import "fmt"

func main() {

	//Aritiméticos
	soma := 2 + 4
	subtracao := 2 - 4
	multitplicacao := 2 * 4
	divisao := 2 / 4

	fmt.Println(soma, subtracao, multitplicacao, divisao)

	var num1 int16 = 3
	var num2 int16 = 4

	// não posso realizar operações com números de tipos diferentes
	total := num1 - num2

	fmt.Println(total)

	//Atribuição
	var variavel string = "teste"
	variavel2 := "teste2"

	fmt.Println(variavel, variavel2)

	//Relacionais
	fmt.Println(1 == 1)
	fmt.Println(1 != 1)
	fmt.Println(1 > 1)
	fmt.Println(1 < 1)
	fmt.Println(1 >= 1)
	fmt.Println(1 <= 1)

	//Lógicos
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	//Unários
	numero := 40
	numero++
	numero += 10
	fmt.Println(numero)

	//Ternário

}
