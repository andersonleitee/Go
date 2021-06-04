package main

import (
	"fmt"
	"os"
)

func main() {
	exibeIntroducao()
	exibeMenu()
	comando := comandoLido()

	switch comando {
	case 0:
		fmt.Println("Sair do programa")
		os.Exit(0) //indicar ao SO para sair com seguranca
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1) //indicar ao SO que foi um erro n esperado
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
