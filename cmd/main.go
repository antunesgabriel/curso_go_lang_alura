package main

import (
	"fmt"
	"os"
	"net/http"
)

func main() {
	name := welcome()
	version := getVersion()


	for {
		option := menu()

		switch option {
		case 1:
			fmt.Println("Iniciando monitoramento...")
			monitoring()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Obrigado", name, "por usar o Monitor em Go Lang - v", version)
			os.Exit(0)

		default:
			fmt.Println("Não entendemos a sua opção, por favor tente novamente...")
		}
	}

}

func welcome() string {
	var name string

	fmt.Println("Olá este é meu primeiro projeto usando Go Lang ❤, por favor me digite seu nome:")
	fmt.Scan(&name)

	fmt.Println("Seja bem vindo", name)
	return name
}

func menu() int {
	var option int
	fmt.Println("Escolha o que deseja fazer:")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Ver logs")
	fmt.Println("0 - Sair do programa")

	fmt.Scan(&option)

	return option
}

func getVersion() float32 {
	return 0.1
}

func monitoring()  {
	urls := []string {
		"https://random-status-code.herokuapp.com",
	}

	urls = append(
		urls,
		"https://www.idcap.org.br/",
		"https://www.fundacaofafipa.org.br/",
		"https://www.fundacaofafipa.org.br/",
	)

	fmt.Println("--- Iniciando o monitoramento de", len(urls), "sites ---")

	for _, url := range urls {
		response, error := http.Get(url)
		statusCode := response.StatusCode

		if error != nil {
			fmt.Println("Ouve um erro inexperado na requisisção ao site", url)
			os.Exit(-1)
		}

		if statusCode == 200 {
			fmt.Println("O site", url, "está respondendo como o esperado. - STATUS CODE:", statusCode)
		} else {
			fmt.Println("O site", url, "não respondeu como esperado. - STATUS CODE", statusCode)
		}
	}
}