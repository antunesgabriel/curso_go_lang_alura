package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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
			readLogs()
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

	file, error := os.Open(filepath.Join("assets", "sites.txt"))

	if error != nil {
		fmt.Println("Erro ao abrir o arquivo", error)
		os.Exit(-1)
	}

	reader := bufio.NewReader(file)

	for {
		url, error := reader.ReadString('\n')

		if error == io.EOF {
			break
		}

		url = strings.TrimSpace(url)

		urls = append(urls, url)
	}

	file.Close()

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
		registerLog(site)
	}
}

func registerLog(site string) {
	file, err := os.OpenFile(filepath.Join("assets", "logs.txt"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)

	if err != nil {
		file.Close()
		fmt.Println("Erro", err)
		return
	}
	timestamps := time.Now().Format("02/01/2006 15:04:05")
	message := "FAILED ás " + timestamps + " -- SITE: " + site + "\n"

	if _, err := file.WriteString(message); err != nil {
		file.Close()
		fmt.Println(err)
		return
	}

	file.Close()
}

func readLogs() {
	if file, err := os.ReadFile(filepath.Join("assets", "logs.txt")); err != nil {
		fmt.Println("Não foi possivel ver os logs")
	} else {
		fmt.Println(string(file))
	}
}