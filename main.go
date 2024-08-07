package main

import (
	"fmt"
	"log"
	"os"
	"simpleCRUD/db"
	"simpleCRUD/models"
)

func main() {
	//Conexao DB
	postgreURL := "postgres://postgres:secret@localhost:5432/simpleCRUD?sslmode=disable"
	db.InitDB(postgreURL)

	//lista opcoes do menu
	showmenu()

	//captura escolha do usuario/inicia loop
	for {
		var escolha int = 1
		fmt.Println("Digite a sua escolha: ")
		// fmt.Scan(&escolha)

		switch escolha {
		default:
			os.Exit(1)
		case 1:
			//criar um novo contatos

			var name string = "Marshall"
			var email string = "Marshall@gmail.com"
			var phone string = "11914105166"

			// fmt.Println("Digite o nome do contato:")
			// fmt.Scanln(&name)
			// fmt.Println("Digite o e-mail:")
			// fmt.Scanln(&email)
			// fmt.Println("Digite o telefone:")
			// fmt.Scanln(&phone)

			contato1 := models.Contact{
				Name:   name,
				Email:  email,
				Phone:  phone,
				Status: true,
			}

			id, err := db.CreateUser(contato1)

			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("O contato foi criado com o ID: %v", id)

		case 2:
			//listar todos os os contatos
		case 3:
			//Filtrar por ID
		case 4:
			//Deletar um registro existente por ID
		case 0:
			//sair do aplicativo
			os.Exit(1)
		}
	}
}

func showmenu() {
	fmt.Println("\n------------------------------------------------------------------")
	fmt.Println("Boas vindas, este é um exemplo de CRUD usando Golang e Postgres!")
	fmt.Println(`Selecione uma opção abaixo:
	1 - Criar um novo registro
	2 - Verificar registros existentes
	3 - Selecionar um determinado registro
	4 - Deletar um registro existente
	0 - Sair da aplicacao`)
}
