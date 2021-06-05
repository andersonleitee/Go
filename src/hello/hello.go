package main

import (
	"fmt"
	"net/http"
	"os"
	"time" //pacote time
)

const monitoramentos = 5
const delay = 10

func main() {
	exibeIntroducao()

	for { //Em Go nao existe While, logo para se realizar um loop usa-se o for (que se usado sem parâmetros roda infinitas vezes até escolher parar)
		exibeMenu()
		comando := comandoLido()

		switch comando {
		case 0:
			fmt.Println("Sair do programa")
			os.Exit(0) //indicar ao SO para sair com seguranca
		case 1:
			monitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1) //indicar ao SO que foi um erro n esperado
		}
	}

}

func exibeIntroducao() {
	nome := "Anderson"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Receber os Logs")
	fmt.Println("0- Sair do monitoramento")
	fmt.Println("")
}

func comandoLido() int { //caso seja uma funcao que retorne algum dado, seu tipo se declara após declaracao da func
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi ", comandoLido)
	fmt.Println("")
	return comandoLido
}

func monitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{"https://www.alura.com.br/", "https://conductor.com.br/", "https://random-status-code.herokuapp.com/"}

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites { //Percorrendo slice com for usando o range
			fmt.Println("Site ", i+1, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second) //declarando a duração do delay
		fmt.Println("")
	}
	fmt.Println("")
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status code:", resp.StatusCode)
	}
}
