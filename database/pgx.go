package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitDB() (*sql.DB, error) {
	connStr := "postgres://postgres:737901@localhost:5432/gowebsocket?sslmode=disable"

	db, err := sql.Open("pgx", connStr)
	if err != nil {

		return nil, fmt.Errorf("opening sql connection %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("checking connection %w", err)
	}

	log.Println("Successfuly connected to DataBase")
	return db, nil
}
