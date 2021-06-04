package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	//exibeIntroducao()
	exibeNomes()

	for { //Em Go nao existe While, logo para se realizar um loop usa-se o for (que se usado sem parâmetros roda infinitas vezes até escolher parar)
		//exibeMenu()
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
}

func comandoLido() int { //caso seja uma funcao que retorne algum dado, seu tipo se declara após declaracao da func
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi ", comandoLido)
	return comandoLido
}

func monitoramento() {
	fmt.Println("Monitorando...")
	var sites [4]string // Array em Go tem que ter tamanho pre definido
	sites[0] = "https://www.alura.com.br/"
	sites[1] = "https://conductor.com.br/"

	fmt.Println(sites)

	site := "https://www.alura.com.br/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status code:", resp.StatusCode)
	}
}

func exibeNomes() { //Ao declarar Slice, n precisa definir seu tamanho
	nomes := []string{"Anderson", "Breno", "Luana", "Catarina"} //Declarando um Slice (tipo abstrato de array) em Go
	fmt.Println(nomes)
	fmt.Println("O tamanho do slice é", len(nomes), "e a capacidade é de", cap(nomes))

	nomes = append(nomes, "Joao")
	//Ao adicionar um item, o Go duplica meu Slice

	fmt.Println(nomes)
	fmt.Println("O tamanho do slice é", len(nomes), "e a capacidade é de", cap(nomes))
}
