package main

import (
	"fmt"
)

func main() {
	//não precisa atribuir a palavra var para variavel basta apenas :=
	nome := "Anderson"
	versao := 1.1
	idade := 24
	fmt.Println("Olá, sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)
}
