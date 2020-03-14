package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeIntroducao()

	for {
		exibeMenu()

		comando := lerComando()

		switch comando {
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		case 1:
			iniciarMonitoramento()
		case 2:
			exibeLogs()
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func exibeIntroducao() {
	nome := "João"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://random-status-code.herokuapp.com/"
	response, _ := http.Get(site)
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		fmt.Println("Site", site, "carregado com sucesso!")
	} else {
		fmt.Println("Site", site, "está com problemas. Status Code:", response.StatusCode)
	}
}

func exibeLogs() {
	fmt.Println("Exibindo logs...")
}
