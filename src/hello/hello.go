package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome = "Anderson" //não precisa dizer o tipo pq o GO faz uma inferência
	var versao float32 = 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("O tipo da varíavel nome é:", reflect.TypeOf(nome))

}
