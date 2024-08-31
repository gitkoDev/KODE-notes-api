package service

import (
	"github.com/gitkoDev/KODE-test-task/internal/repository"
	"github.com/gitkoDev/KODE-test-task/models"
)

type NotesService struct {
	notesRepository repository.Notes
}

func NewNotesService(notesRepository repository.Notes) *NotesService {
	return &NotesService{
		notesRepository: notesRepository,
	}
}

func (s *NotesService) AddNote(userId int, note models.Note) (int, error) {
	return s.notesRepository.AddNote(userId, note)
}

func (s *NotesService) GetAllNotes(userId int) ([]models.Note, error) {
	return s.notesRepository.GetAllNotes(userId)
}
