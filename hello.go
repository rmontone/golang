package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao()

	exibeMenu()

	comando := leComando()

	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		exibeLogs()
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Não conheço esse comando")
		os.Exit(-1)
	}
}

func exibeIntroducao() {
	nome := "Henrique"
	version := 1.8
	fmt.Println("Hello, Mr.", nome)
	fmt.Println("This program is in the", version, "version")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa")

}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://random-status-code.herokuapp.com"
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, " foi carregado com sucesso!")
	} else {
		fmt.Println("Site: ", site, "está com problemas.\nStatus Code: ", resp.StatusCode)
	}
}

func exibeLogs() {
	fmt.Println("Exibindo Logs...")

}
