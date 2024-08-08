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

	//captura escolha do usuario/inicia loop
	for {
		//lista opcoes do menu
		escolha := showmenu()

		switch escolha {
		default:
			os.Exit(1)
		case 1:
			//criar um novo contatos

			var name string
			var email string
			var phone string

			fmt.Println("Digite o nome do contato:")
			fmt.Scanln(&name)
			fmt.Println("Digite o e-mail:")
			fmt.Scanln(&email)
			fmt.Println("Digite o telefone:")
			fmt.Scanln(&phone)

			contato1 := models.Contact{
				Name:   name,
				Email:  email,
				Phone:  phone,
				Status: true,
			}

			id, err := db.CreateContact(contato1)

			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("O contato foi criado com o ID: %v", id)

		case 2:
			contatos, err := db.ListContacts()
			if err != nil {
				log.Fatal(err)
			}
			for _, v := range contatos {
				fmt.Printf("ID: %v | Name: %v | Email: %v | Phone: %v \n", v.ID, v.Name, v.Email, v.Phone)

			}
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

func showmenu() int {
	fmt.Println("\n------------------------------------------------------------------")
	fmt.Println("Boas vindas, este é um exemplo de CRUD usando Golang e Postgres!")
	fmt.Println(`Selecione uma opção abaixo:
	1 - Criar um novo registro
	2 - Verificar registros existentes
	3 - Selecionar um determinado registro
	4 - Deletar um registro existente
	0 - Sair da aplicacao`)

	var escolha int
	fmt.Println("\n Digite a sua escolha: ")
	fmt.Scan(&escolha)
	return escolha

}
