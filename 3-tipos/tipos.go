package main

import (
	"errors"
	"fmt"
)

func main() {

	var num1 int8 = 10
	fmt.Println(num1)

	var num2 uint32 = 1000 // int sem sinal
	fmt.Println(num2)

	//alias
	var num3 rune = 1111
	fmt.Println(num3)

	var num4 byte = 116
	fmt.Println(num4)

	var num5 float32 = 12.5
	fmt.Println(num5)

	var num6 float32 = 122.5
	fmt.Println(num6)

	//string
	var str string = "Teste"
	fmt.Println(str)

	str2 := "rr"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	// todo tipo de dado tem seu valor inicial
	var text int16
	fmt.Println(text)

	var text2 string
	fmt.Println(text2)

	//boleano
	var boleano bool = true
	fmt.Println(boleano)

	//error
	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("vix deu erro")
	fmt.Println(erro2)

}
