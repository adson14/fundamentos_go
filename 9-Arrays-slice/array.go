package main

import "fmt"

func main() {
	fmt.Println("Arrays e slice")

	//Decaração do array
	var array1 [5]string
	array1[0] = "Primeira posição"
	fmt.Println(array1)

	// Declaração com inferência de tipos
	var array2 = [5]string{
		"Segunda posição",
		"Terceira posição",
		"Quarta posição",
		"Quinta posição",
		"Sexta posição",
	}

	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//Slice
	slice := []int{1, 2, 23, 333, 2}
	fmt.Println(slice)

	//Adiciona mais um item e retorna um novo slice
	slice = append(slice, 55)
	fmt.Println(slice)

	//Arrays Internos
	slice2 := make([]float32, 10, 11)
	fmt.Println(slice2)

	/*
		Ao ultrapassar o tamanho do array, um novo array é criado dobrando o tamanho
	*/
	slice2 = append(slice2, 1.1, 2.2, 3.3, 4.4)
	fmt.Println(slice2)

}
