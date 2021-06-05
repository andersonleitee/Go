package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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
			imprimeLogs()
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1) //indicar ao SO que foi um erro n esperado
		}
	}

}

func exibeIntroducao() {
	fmt.Print("Digite seu nome:")
	var nome string
	fmt.Scan(&nome)
	versao := 1.1
	fmt.Println("")
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("")
}

func exibeMenu() {
	fmt.Println("******************************")
	fmt.Println("*          Menu              *")
	fmt.Println("******************************")
	fmt.Println("")
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

	sites := leSitesDoArquivo()

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
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		criaLog(site, true)
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status code:", resp.StatusCode)
		criaLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
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

func criaLog(site string, status bool) {
	arquivo, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- Online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()

}

func imprimeLogs() {

	fmt.Println("Exibindo Logs...")
	arquivo, err := ioutil.ReadFile("logs.txt") //ioutil serve como alternativa p/ abrir arquivo e n precisa fechar.

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	fmt.Println(string(arquivo))

}
