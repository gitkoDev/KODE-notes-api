package service

import (
	"github.com/gitkoDev/KODE-test-task/internal/repository"
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
	GenerateToken(userName string, userPassword string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Service struct {
	Auth
	Notes
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Auth:  NewAuthService(repository.Auth),
		Notes: NewNotesService(repository.Notes),
	}
}
