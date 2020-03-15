package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

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
	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')

		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func testaSite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	if response.StatusCode >= 200 && response.StatusCode < 300 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", response.StatusCode)
	}
}

func exibeLogs() {
	fmt.Println("Exibindo logs...")
}
