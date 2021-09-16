package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string
	var age float32
	version := 1.1 // declarando variavel e atribuindo um valor

	fmt.Println("-> Endereço de idade", &age)
	fmt.Println("-> Endereço de nome", &name)

	fmt.Println("- Olá, tudo bem? Por favor digite seu nome:")
	fmt.Scan(&name)

	fmt.Println("- Prazer", name, "digite agora sua idade:")
	fmt.Scan(&age)

	fmt.Println("- Seu nome é", name, " e sua idade é ", age, "anos")

	fmt.Println("--- Hello Word em Go v", version, "---")
	fmt.Println("Tipo inferido a version: ", reflect.TypeOf(version))
}