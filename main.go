package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./clientes.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS clientes (id INTEGER PRIMARY KEY, nome TEXT, email TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Banco de dados inicializado com sucesso!")
}
