package db

import (
	"database/sql"
	"fmt"
	"log"
	"simpleCRUD/models"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(port int, host, user, password, dbname string) {
	var err error
	dataSourceName := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("database connected successfully!")
}

func CreateContact(u models.Contact) (int, error) {
	var id int
	query := "INSERT INTO contato (name, email, phone, status) VALUES ($1, $2, $3, $4) returning id "
	err := db.QueryRow(query, u.Name, u.Email, u.Phone, u.Status).Scan(&id)

	return id, err

}

func ListContacts() ([]models.Contact, error) {
	query := `SELECT id, name, email, phone FROM contato`
	rows, err := db.Query(query)

	var contatos []models.Contact

	for rows.Next() {
		var id int
		var name string
		var email string
		var phone string

		rows.Scan(&id, &name, &email, &phone)
		c1 := models.Contact{
			ID:    id,
			Name:  name,
			Email: email,
			Phone: phone,
		}
		contatos = append(contatos, c1)

		// fmt.Printf("ID: %v | Name: %v | Email: %v | Phone: %v \n", id, name, email, phone)
	}

	return contatos, err

}

func ContatDetails(id int) (models.Contact, error) {
	query := `SELECT * FROM contato WHERE ID=$1`
	rows := db.QueryRow(query, id)

	var name string
	var email string
	var phone string
	var status bool

	err := rows.Scan(&id, &name, &email, &phone, &status)
	c1 := models.Contact{
		ID:     id,
		Name:   name,
		Email:  email,
		Phone:  phone,
		Status: status,
	}

	return c1, err

}

func DeleteContact(id int) (sql.Result, error) {
	query := `DELETE FROM contato WHERE ID=$1`

	result, err := db.Exec(query, id)

	if err != nil {
		log.Println(err)
	}

	return result, err

}
