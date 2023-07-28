package main

import (
	"fmt"
)

func main() {
	nome := "Henrique"
	version := 1.8
	fmt.Println("Hello, Mr.", nome)
	fmt.Println("This program is in the", version, "version")

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O endereço da minha variável comando é", &comando)
	fmt.Println("O comando escolhido foi", comando)
}
