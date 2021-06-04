package main

import (
	"fmt"
)

func main() {
	nome := "Anderson"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Receber os Logs")
	fmt.Println("0- Sair do monitoramento")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi ", comando)

	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo Logs...")
	} else if comando == 0 {
		fmt.Println("Sair do programa")
	} else {
		fmt.Println("Não conheço este comando")
	}

	//No Go, além de n precisar de parenteses, a clausula if apenas será efetiva se retornar um valor booleano
	maiordeidade := true

	if maiordeidade {
		fmt.Println("ok")
	}

}
