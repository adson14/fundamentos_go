package main

import "fmt"

// Ponteiro é uma referência de memória
func main() {
	var1 := 10
	var2 := var1

	var1++

	fmt.Println(var1, var2)

	var3 := 100
	var ponteiro *int
	ponteiro = &var3

	fmt.Println(*ponteiro)

	var3++

	fmt.Println(*ponteiro)
}
