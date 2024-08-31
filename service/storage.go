package service

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gitkoDev/KODE-test-task/helpers"
	"github.com/gitkoDev/KODE-test-task/models"
	"github.com/golang-jwt/jwt/v5"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db}
}

func (s *Storage) CreateUser(user models.User) (int, error) {
	id, err := s.CheckForUserExistence(user.Name, user.Password)

	if err != nil {
		return -1, err
	}

	if id != -1 {
		errorResponse := fmt.Sprintf("user %s already exists in database", user.Name)
		return 0, errors.New(errorResponse)
	}

	var user_id int
	user.Password = helpers.GeneratePasswordHash(user.Password)

	query := `INSERT INTO users (name, password_hash) VALUES ($1, $2) RETURNING id`
	if err := s.db.QueryRow(query, user.Name, user.Password).Scan(&user_id); err != nil {
		return user_id, err
	}

	return user_id, nil
}

func (s *Storage) GetUser(name string, password string) (models.User, error) {
	id, err := s.CheckForUserExistence(name, password)

	if err != nil {
		return models.User{}, err
	}

	if id == -1 {
		return models.User{}, err
	}

	return models.User{Id: id, Name: name, Password: password}, nil

}

func (s *Storage) GetUserById(id int) (int, error) {
	query := `SELECT 1 FROM users WHERE id = $1`

	var user_id int
	err := s.db.QueryRow(query, id).Scan(&user_id)
	if err != nil {
		return -1, err
	}

	return user_id, nil
}

func (s *Storage) AddNote(userId int, note models.Note) (int, error) {
	_, err := s.GetUserById(userId)
	if err != nil {
		return -1, err
	}
	var id int

	note.Validate()

	query := `INSERT INTO notes (user_id, content) VALUES ($1, $2) RETURNING id`
	if err := s.db.QueryRow(query, userId, note.Content).Scan(&id); err != nil {
		return -1, err
	}

	return id, nil

}

func (s *Storage) GetAllNotes(userId int) ([]models.Note, error) {
	// user, err := s.GetUserById(userId)
	// if err != nil {
	// 	return []models.Note{}, err
	// }
	// var id int

	query := `SELECT content FROM notes WHERE user_id = $1`
	res, err := s.db.Query(query, userId)
	if err != nil {
		return []models.Note{}, err
	}
	defer res.Close()

	userNotes := []models.Note{}

	for res.Next() {
		note := models.Note{}

		err := res.Scan(&note.Content)
		if err != nil {
			return []models.Note{}, nil
		}

		userNotes = append(userNotes, note)
	}

	return userNotes, nil
}

func (s *Storage) CheckForUserExistence(name string, password string) (int, error) {
	passwordHash := helpers.GeneratePasswordHash(password)
	query := `SELECT id FROM users WHERE name = $1 AND password_hash = $2`
	var id int

	if err := s.db.QueryRow(query, name, passwordHash).Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return -1, nil
		}
		return -1, err
	}

	return id, nil

}

func (s *Storage) GenerateToken(name string, password string) (string, error) {
	user, err := s.GetUser(name, password)
	if err != nil {
		return "", err
	}

	claims := models.UserTokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
	if err != nil {
		return "", err
	}
	return ss, nil
}

func (s *Storage) ParseToken(accessToken string) (int, error) {

	token, err := jwt.ParseWithClaims(accessToken, &models.UserTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNING_KEY")), nil
	})

	if err != nil {
		return 0, err
	} else if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return 0, errors.New("wrong token signing method")
	} else if claims, ok := token.Claims.(*models.UserTokenClaims); !ok {
		return 0, errors.New("wrong token type")
	} else {
		return claims.Id, nil
	}

}
