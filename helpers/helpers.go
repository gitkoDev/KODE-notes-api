package helpers

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gitkoDev/KODE-test-task/models"
	"github.com/sirupsen/logrus"
)

type ErrResponseJSON struct {
	ErrMsg string `json:"error"`
}

func DecodeAuthJSON(req *http.Request) (models.User, error) {
	var input models.User

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&input); err != nil {
		return input, err
	}
	return input, nil
}

func DecodeNotesJSON(req *http.Request) (models.Note, error) {
	var input models.Note

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&input); err != nil {
		return input, err
	}
	return input, nil
}

func RespondWithJSON(w http.ResponseWriter, data any, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		return err
	}
	return nil
}

func RespondWithError(w http.ResponseWriter, receivedErr error, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	var errMessage ErrResponseJSON

	if receivedErr.Error() == "EOF" {
		errMessage = ErrResponseJSON{ErrMsg: "please provide valid input data"}
	} else {
		errMessage = ErrResponseJSON{ErrMsg: receivedErr.Error()}

	}

	err := json.NewEncoder(w).Encode(errMessage)
	if err != nil {
		logrus.Fatal("error encoding json:", err)
	}

}

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("HASH_SALT"))))
}
