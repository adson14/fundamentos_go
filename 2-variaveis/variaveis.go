package main

import "fmt"

func main() {

	var variavel1 string = "Explicito"
	fmt.Println(variavel1)

	//Inferencia de tipos
	variavel2 := "implicito"

	fmt.Println(variavel2)

	var (
		var1 string = "a"
		var2 int    = 12
	)

	fmt.Println(var1, var2)

	v1, v2 := "1", "2"
	fmt.Println(v1, v2)

	v1, v2 = v2, v1

	fmt.Println(v1, v2)

	// Constantes
	const constant1 string = "S"
	fmt.Println(constant1)

}
