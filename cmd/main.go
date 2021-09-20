package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const REPEAT_AMOUNT int = 5
const WAIT time.Duration = 10

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

	fmt.Println("")

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

	for count := 0; count < REPEAT_AMOUNT; count++  {
		for _, url := range urls {
			testSite(url)
		}
		fmt.Println("")

		time.Sleep(WAIT * time.Minute)
	}
}

func testSite(site string) {
	response, error := http.Get(site)
	statusCode := response.StatusCode

	if error != nil {
		fmt.Println("OPS --  Ouve um erro durante o processo")
		os.Exit(-1)
	}

	if statusCode == 200 {
		fmt.Println("SUCCESS -- O site:", site, "está respondendo como o esperado - STATUS CODE", statusCode)
	} else {
		fmt.Println("FAILED -- O site:", site, "não respondeu como o esperado - STATUS CODE", statusCode)
	}
}