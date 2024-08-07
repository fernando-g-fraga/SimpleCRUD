package db

import (
	"database/sql"
	"fmt"
	"log"
	"simpleCRUD/models"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("database connected successfully!")
}

func CreateUser(u models.Contact) (int, error) {
	var id int
	query := "INSERT INTO contato (name, email, phone, status) VALUES ($1, $2, $3, $4) returning id "
	err := db.QueryRow(query, u.Name, u.Email, u.Phone, u.Status).Scan(&id)

	return id, err

}
