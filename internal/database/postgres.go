package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

type DataBase struct {
	sql *sql.DB
}

func NewDB() *DataBase {
	return &DataBase{}
}

func (d *DataBase) Connect(c context.Context) (*sql.DB, error) {
	ctx, cancel := context.WithCancel(c)
	defer cancel()

	str := fmt.Sprintf("postgresql://%s", os.Getenv("DB"))
	db, err := sql.Open("postgres", str)
	log.Println(str)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err := RetryPing(db); err != nil {
		log.Println(err)
		return nil, err
	}

	schema := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(20),
		surname VARCHAR(30),
		age SMALLINT,
		email VARCHAR(30) NOT NULL,
		phone_number VARCHAR(20),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	if _, err := db.ExecContext(ctx, schema); err != nil {
		panic("Error creating migrations: " + err.Error())
	}

	d.sql = db

	log.Println("Database connection is valid")
	return db, nil
}

func RetryPing(db *sql.DB) error {
	var err error
	for i := 0; i < 5; i++ {
		if err = db.Ping(); err == nil {
			return nil
		}
		time.Sleep(time.Second)
	}
	return nil
}

func (d *DataBase) CloseDB() error {
	return d.sql.Close()
}
