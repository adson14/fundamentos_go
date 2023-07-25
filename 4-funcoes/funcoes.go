package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func calculosMatematicos(num1, num2 int8) (int8, int8) {
	soma := num1 + num2
	sub := num1 - num2

	return soma, sub
}

func main() {
	soma := somar(10, 30)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("texto")
	fmt.Println(resultado)

	resCalculoSoma, _ := calculosMatematicos(22, 14)

	fmt.Println(resCalculoSoma)

}
