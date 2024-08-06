/*
1  - Boas vindas, este é um exemplo de CRUD usando Golang e Postgres
2 - Selecione uma opcao abaixo:
	1 - Criar um novo registro
	2 - Verificar registros existentes
	3 - Selecionar um determinado registro
	4 - Deletar um registro existente
	0 - Sair da aplicacao
Desenvolvido por Fernando Fraga para fins de portfólio.
*/

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

// type contato struct {
// 	name   string
// 	email  string
// 	phone  string
// 	status bool
// }

func main() {
	// Conecao do banco de dados
	db := ConectarDB()
	CreateTable(db)

	err := db.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for {
		//menu e escolha do usuário
		showmenu()
		var escolha int
		fmt.Println("Digite a sua escolha: ")
		fmt.Scan(&escolha)

		switch escolha {
		default:
			os.Exit(1)
		case 1:
			createRecord(db)
		case 2:
			selectAll(db)
		case 3:
			var id int
			fmt.Println("Digite o código desejado: ")
			fmt.Scan(&id)
			selectOne(db, id)
		case 4:
			var id int
			var hold string
			fmt.Println("Digite o código desejado: ")
			fmt.Scan(&id)
			fmt.Printf("Deseja realmente seguir com a exclusão do id: %v ? y or n", id)
			fmt.Scan(&hold)
			deleteRecord(db, id)
		case 5:
			os.Exit(1)
		}
	}
}

func createRecord(db *pgx.Conn) {

	var first string
	var email string
	var phone string
	status := true

	fmt.Println("Digite o nome do seu contato: ")
	fmt.Scan(&first)
	fmt.Println("Digite o email: ")
	fmt.Scan(&email)
	fmt.Println("Digite o telefone: ")
	fmt.Scan(&phone)

	query := `INSERT INTO contato (name,email,phone,status)
	VALUES ($1,$2,$3,$4) returning ID`
	_, err := db.Query(context.Background(), query, first, email, phone, status)

	if err != nil {
		log.Fatal("Ocorreu um erro ao cadastrar o usuário ", err)
	}

	defer db.Close(context.Background())

}
func selectAll(db *pgx.Conn) {
	query := "SELECT id,name,email,phone,status FROM contato"
	rows, err := db.Query(context.Background(), query)

	if err != nil {
		log.Fatal("Ocorreu um erro ao consultar os dados ", err)
	}

	var id int
	var name string
	var email string
	var phone string
	var status bool

	for rows.Next() {
		err = rows.Scan(&id, &name, &email, &phone, &status)
		if err != nil {
			log.Fatal("Ocorreu um erro ao consultar os dados ", err)
		}
		fmt.Printf("ID: %v | Name: %v | Email: %v | Phone: %v | Status: %v \n", id, name, email, phone, status)
	}

	fmt.Println("Digite alguma tecla para retornar ao menu anterior.")
	var hold string
	fmt.Scan(&hold)

	defer db.Close(context.Background())

}
func selectOne(db *pgx.Conn, id int) {

	var name string
	var email string
	var phone string
	var status bool

	query := "SELECT name,email,phone,status FROM contato WHERE ID=$1"
	db.QueryRow(context.Background(), query, id).Scan(&name, &email, &phone, &status)

	fmt.Printf("ID: %v | Name: %v | Email: %v | Phone: %v | Status: %v \n", id, name, email, phone, status)

}
func deleteRecord(db *pgx.Conn, id int) {
	query := "DELETE FROM contato WHERE id=$1"
	_, err := db.Query(context.Background(), query, id)

	if err != nil {
		log.Fatal("Ocorreu um erro ao consultar os dados ", err)
	} else {
		fmt.Println("Registro excluído com sucesso!")
	}

}

func showmenu() {
	fmt.Println("\n -----------------------------------------------------------------")
	fmt.Println("Boas vindas, este é um exemplo de CRUD usando Golang e Postgres!")
	fmt.Println(`Selecione uma opção abaixo:
	1 - Criar um novo registro
	2 - Verificar registros existentes
	3 - Selecionar um determinado registro
	4 - Deletar um registro existente
	0 - Sair da aplicacao`)
}

func ConectarDB() *pgx.Conn {
	postgreURL := "postgres://postgres:secret@localhost:5432/simpleCRUD"
	db, err := pgx.Connect(context.Background(), postgreURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return db

}

func CreateTable(db *pgx.Conn) {
	query := `CREATE TABLE IF NOT EXISTS contato (
	id SERIAL PRIMARY KEY,
	name VARCHAR(20),
	email VARCHAR(20),
	phone VARCHAR(20),
	status BOOLEAN
	)`

	_, err := db.Exec(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
}
