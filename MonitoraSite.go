package main

import "fmt"
import "os"
import "net/http"

func main() {

  diferente()

 for {

  menuchefe()
  
  comando := comandopai()

  //fmt.Println("**o endereço da variável comando é:", &comando)

  switch comando {
     
     case 1: 
    
            IM()
    
     case 2: 
           
            fmt.Println("\nExibindo logs...")       
     
     case 0: 
           
            fmt.Println("\nSaindo do programa")
            os.Exit(0)
     
     default:
           
            fmt.Println("\nNão conheço este comando digitado")  
            os.Exit(-1)           
  
  } 
 }
}

func diferente() {

  var nome string
  versao := 1.3
 
  fmt.Println("Olá usuário! Digite teu nome: ")
  fmt.Scan(&nome)

  fmt.Println("Olá senhor(a)", nome)
  fmt.Println("\nPara fins de curiosidade, este programa está na versão", versao)

}

func menuchefe() {

  fmt.Println("\nMenu de tarefas:") 
  fmt.Println("1- Inciar Monitoramento")
  fmt.Println("2- Exibir logs")
  fmt.Println("0- Sair do programa")

}

func comandopai() int {

  var comando int

  fmt.Scan(&comando)

  fmt.Println("\nVocê escolheu o comando: ", comando)

  return comando

}

func IM() {
     
  var site string 
  fmt.Println("\nDigite o site que você quer monitorar: ")
  fmt.Scan(&site)

  fmt.Println("\nMonitorando a aplicação...")
  
  resposta, _ := http.Get(site)
  //fmt.Println(resposta)

  if resposta.StatusCode == 200{

    fmt.Println("O site: ", site, "foi carregado com sucesso!")

  }else{

    fmt.Println("O site: ", site, "está com problemas. StatusCode", resposta.StatusCode)

  }

}

