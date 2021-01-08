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

	type User struct {
		id, age                    int
		firstName, lastName, email string
	}

	var users []User
	// var id, age int
	// var firstName, lastName, email string

	// err = db.QueryRow(`
	// 	INSERT INTO users(age, first_name, last_name, email)
	// 	VALUES ($1, $2, $3, $4) RETURNING id`,
	// 	36, "Maria", "Geo", "mariageo@gmail.com",
	// ).Scan(&id)

	// err = db.QueryRow(`
	// SELECT id, first_name, email FROM users WHERE id=$1`, 3).Scan(&id, &name, &email)

	rows, err := db.Query(`
	SELECT id, age, first_name, last_name, email FROM users WHERE age > $1`, 30)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		if err = rows.Scan(&user.id, &user.age, &user.firstName, &user.lastName, &user.email); err != nil {
			panic(err)
		}
		// fmt.Println("id:", id, "Age:", age, "First name:", firstName, "Last name:", lastName, "Email:", email)
		users = append(users, user)
		fmt.Println(users)
	}
}
