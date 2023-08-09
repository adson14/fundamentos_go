package main

import "fmt"

func main() {
	fmt.Println("maps")

	usuario := map[string]string{
		"nome":      "Adson",
		"sobrenome": "Silva",
		"idade":     "20",
	}

	fmt.Println(usuario["nome"])

	usuario["tel"] = "9999-9999"

	fmt.Println(usuario)
}
