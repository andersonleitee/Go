package main

import "fmt"

//para declarar um struct usa-se o type antes do nome da estrutura e "struct" após
type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	contaAnderson := ContaCorrente{titular: "Anderson", numeroAgencia: 589, numeroConta: 123456, saldo: 150}
	contaBreno := ContaCorrente{"Breno", 789, 7894561, 250}

	//outra forma de usar struct
	var contaJulia *ContaCorrente
	contaJulia = new(ContaCorrente)
	contaJulia.titular = "Julia"
	fmt.Println(contaAnderson, contaBreno, *contaJulia)
}
