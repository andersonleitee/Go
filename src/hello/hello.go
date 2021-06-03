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
	fmt.Println(comando)

}
