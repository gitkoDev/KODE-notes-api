package repository

import (
	"database/sql"

	"github.com/gitkoDev/KODE-test-task/models"
)

type NotesPostgres struct {
	db *sql.DB
}

func NewNotesPostgres(db *sql.DB) *NotesPostgres {
	return &NotesPostgres{db: db}
}

func (s *NotesPostgres) AddNote(userId int, note models.Note) (noteId int, err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return
	}

	// Yandex.Speller validation
	note.Validate()

	query := `INSERT INTO notes (user_id, content) VALUES ($1, $2) RETURNING id`
	if err := s.db.QueryRow(query, userId, note.Content).Scan(&noteId); err != nil {
		tx.Rollback()
		return 0, err
	}

	return noteId, tx.Commit()
}

func (s *NotesPostgres) GetAllNotes(userId int) ([]models.Note, error) {
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
