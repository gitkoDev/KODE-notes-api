package repository

import (
	"database/sql"
	"fmt"

	"github.com/gitkoDev/KODE-test-task/helpers"
	"github.com/gitkoDev/KODE-test-task/models"
	"github.com/sirupsen/logrus"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (s *AuthPostgres) CreateUser(user models.User) (userId int, err error) {
	id, err := s.CheckForUserExistence(user.Name, user.Password)

	if err != nil && err != sql.ErrNoRows {
		logrus.Println("returning with error:", err)
		return
	}

	if id != 0 {
		responseString := fmt.Errorf("user %s exists in database", user.Name)
		return 0, responseString
	}

	var user_id int
	user.Password = helpers.GeneratePasswordHash(user.Password)

	query := `INSERT INTO users (name, password_hash) VALUES ($1, $2) RETURNING id`
	if err := s.db.QueryRow(query, user.Name, user.Password).Scan(&user_id); err != nil {
		return user_id, err
	}

	return user_id, nil
}

func (s *AuthPostgres) GetUser(name string, password string) (user models.User, err error) {
	id, err := s.CheckForUserExistence(name, password)

	if err != nil {
		return
	}

	if id == 0 {
		return
	}

	return models.User{Id: id, Name: name, Password: password}, nil

}

func (s *AuthPostgres) CheckForUserExistence(name string, password string) (userId int, err error) {
	passwordHash := helpers.GeneratePasswordHash(password)
	query := `SELECT id FROM users WHERE name = $1 AND password_hash = $2`
	var id int

	if err := s.db.QueryRow(query, name, passwordHash).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil

}
