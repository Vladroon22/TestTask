package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"gthub.com/Vladroon22/TestTask/internal/entity"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repo {
	return Repo{
		db: db,
	}
}

func (rp Repo) CreateUser(c context.Context, user entity.User) error {
	ctx, cancel := context.WithCancel(c)
	defer cancel()

	query := "INSERT INTO users (name, surname, email, phone_number, age) VALUES ($1, $2, $3, $4, $5)"
	if _, err := rp.db.ExecContext(ctx, query, user.FirstName, user.LastName, user.Email, user.Phone, &user.Age); err != nil {
		log.Println(err)
		return errors.New("bad response from database")
	}

	log.Println("User successfully added")
	return nil
}

func (rp Repo) UpdateUser(c context.Context, user entity.User) (int, error) {
	ctx, cancel := context.WithCancel(c)
	defer cancel()

	query2 := "UPDATE users SET name = $1, surname = $2, email = $3, phone_number = $4, age = $5 WHERE id = $6"
	if _, err := rp.db.ExecContext(ctx, query2, user.FirstName, user.LastName, user.Email, user.Phone, &user.Age, user.ID); err != nil {
		log.Println(err)
		return 0, errors.New("bad response from database")
	}

	log.Println("User successfully updated")
	return user.ID, nil
}

func (rp Repo) GetUser(c context.Context, id int) (entity.User, error) {
	ctx, cancel := context.WithCancel(c)
	defer cancel()

	user := entity.User{}
	query := "SELECT id, name, surname, email, phone_number, age FROM users WHERE id = $1"
	err := rp.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.Age)
	if err != nil {
		log.Println(err)
		return entity.User{}, errors.New("bad response from database")
	}
	return user, nil
}
