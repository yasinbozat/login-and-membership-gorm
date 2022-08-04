package main

import (
	"database/sql"
	"fmt"
	"log"

	//"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "db_user"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT nickname FROM tb_user WHERE id=1")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", name)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	/*
		app := fiber.New()

		app.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("")
		})

		app.Listen(":3000")*/
}
