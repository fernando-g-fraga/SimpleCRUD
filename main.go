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

type contato struct {
	name   string
	email  string
	phone  string
	status bool
}

func main() {

	db := ConectarDB()
	CreateTable(db)

	err := db.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	escolha := showmenu()

	switch escolha {
	default:
		os.Exit(1)
	case 1:
		rows := createRecord(db)
		fmt.Println(rows)
	case 2:
		selectAll()
	case 3:
		selectOne()
	case 4:
		deleteRecord()
	case 5:
		os.Exit(1)
	}

}

func createRecord(db *pgx.Conn) pgx.Rows {

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
	rows, err := db.Query(context.Background(), query, first, email, phone, status)

	if err != nil {
		log.Fatal("Ocorreu um erro ao cadastrar o usuário ", err)
	}

	return rows

}
func selectAll()    {}
func selectOne()    {}
func deleteRecord() {}

func showmenu() int {
	fmt.Println("Boas vindas, este é um exemplo de CRUD usando Golang e Postgres!")
	fmt.Println(`Selecione uma opção abaixo:
	1 - Criar um novo registro
	2 - Verificar registros existentes
	3 - Selecionar um determinado registro
	4 - Deletar um registro existente
	0 - Sair da aplicacao`)

	var escolha int

	fmt.Scan(&escolha)

	return escolha

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
