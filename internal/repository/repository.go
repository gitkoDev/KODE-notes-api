package repository

import (
	"database/sql"

	"github.com/gitkoDev/KODE-test-task/models"
)

type Notes interface {
	AddNote(userId int, note models.Note) (int, error)
	GetAllNotes(userId int) ([]models.Note, error)
}

type Auth interface {
	CreateUser(user models.User) (int, error)
	GetUser(userName string, userPassword string) (models.User, error)
	CheckForUserExistence(userName string, userPassword string) (int, error)
}

type Repository struct {
	Notes
	Auth
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Auth:  NewAuthPostgres(db),
		Notes: NewNotesPostgres(db),
	}
}
