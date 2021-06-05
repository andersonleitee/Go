
<p> <img src="https://cdn.worldvectorlogo.com/logos/gopher.svg" width="80">  <img src="https://raw.githubusercontent.com/geiltonxavier/aprenda-go/master/gopher.png" width="80">  <img src="https://blog.golang.org/gopher/gopher.png" width="80"> 

# Golang
Go, como popularmente é conhecida, é uma linguagem de programação criada pela Google e lançada em código livre em novembro de 2009. É uma linguagem compilada e focada em produtividade e programação concorrente. 

O repositório em questão será utilizado para praticar e estudar a linguagem. Em paralelo, uma aplicação utilanzando 100% Golang será criada como forma de solidificar os conceitos aprendidos no curso de Golang da Alura. 
  
# Aplicação
O programa tem como objetivo monitorar sites e verificar se estão ativos ou não conforme a URL passada no arquivo `sites.tx`.

![](https://github.com/andersonleitee/Go/blob/main/prints/Arquivo_sites.png?raw=true) 
|:--:| 
| *Figura 1: sites.tx, arquivo.* |

### Menu
Aqui, será solicitado ao usuário que digite o seu nome para que seja montada a introdução da aplicação com o nome e versão do programinha. Em seguida será disposto o Menu de opções em que o usuário escolherá o número que seja condizente.

![](https://github.com/andersonleitee/Go/blob/main/prints/Menu.png?raw=true) 
|:--:| 
| *Figura 2: Menu.* |  

### Monitoramento

Nesta seção, caso o usuário tenha escolhido o número 1, será realizado o monitoramento de fato dos sites que estão no arquivo `sites.txt`. A função `monitoramento()` e `testaSite(site string)` são cruciais para esta etapa, no qual por meio da requisição GET será possível identificar se o site está ativo ou inativo (devido a algum erro).

![](https://github.com/andersonleitee/Go/blob/main/prints/Monitorando.png?raw=true) 
|:--:| 
| *Figura 3: Monitoramento.* | 

### Exibição de Logs
Caso o usuário tenha escolhido a opção 2, será exibido os Logs do monitoramento. Esses Logs estão no arquivo `logs.txt`, o qual é escrito a medido que é realizado o moniroramento.

![](https://github.com/andersonleitee/Go/blob/main/prints/Arquivo_Logs.png?raw=true) 
|:--:| 
| *Figura 4: logs.txt, arquivo.* |  

![](https://github.com/andersonleitee/Go/blob/main/prints/Exibindo_logs.png?raw=true) 
|:--:| 
| *Figura 5: Exibição dos Logs.* |

### Saída
Por fim, caso o usuário tenha escolhido a opção 0, sairá do do programa.

![](https://github.com/andersonleitee/Go/blob/main/prints/Saida.png?raw=true) 
|:--:| 
| *Figura 6: Saída.* |

  
