package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "darkroom_dev"
)

func main() {
	// %s for string, %d for int
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var id int
	err = db.QueryRow(`
		INSERT INTO users(age, first_name, last_name, email)
		VALUES ($1, $2, $3, $4) RETURNING id`,
		36, "Maria", "Geo", "mariageo@gmail.com",
	).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("id:", id)
}
