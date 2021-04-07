package main

import (
	"fmt"
	"log"
	"os"
)

//basic registration system

func main() {
	f, err := os.Create("dados.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	for i := 1; i > 0; i++ {
		var (
			nome  string
			idade string
		)

		fmt.Println("Seu nome:")
		fmt.Scanln(&nome)
		fmt.Println("Sua idade:")
		fmt.Scanln(&idade)

		_, err2 := f.WriteString("\n Nome: " + nome + " Idade: " + idade)

		if err2 != nil {
			log.Fatal(err2)
		}
	}
}
