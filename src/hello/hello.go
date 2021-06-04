package main

import (
	"fmt"
	//"net/http"
	"os"
)

func main() {
	// exibeIntroducao()
	// exibeMenu()

	//nome, idade := nomeEIdade() //declaracao de funcao com mais de uma variavel
	nome, _ := nomeEIdade() //caso deseje ignorar um dos dados retornados, o Go identifica o underline '_' como dado vazio

	fmt.Println(nome)

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

//Funcoes que retornam mais de um dado
//1 - Devem ser daclarados os tipos de variáveis que serão retornados
//2 - A declaração dos tipos devem estar entre parenteses
//3 - retorne na ordem que declarou os tipos de retorno

func nomeEIdade() (string, int) {
	nome := "Anderson"
	idade := 24

	return nome, idade
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
	//site := "https://www.alura.com.br/"
	//resp, err := http.Get(site)
}
