package service

import (
	"errors"
	"os"
	"time"

	"github.com/gitkoDev/KODE-test-task/internal/repository"
	"github.com/gitkoDev/KODE-test-task/models"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	authRepository repository.Auth
}

func NewAuthService(authRepository repository.Auth) *AuthService {
	return &AuthService{authRepository: authRepository}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	return s.authRepository.CreateUser(user)
}

func (s *AuthService) GetUser(userName string, userPassword string) (models.User, error) {
	return s.authRepository.GetUser(userName, userPassword)
}

func (s *AuthService) CheckForUserExistence(userName string, userPassword string) (int, error) {
	return s.authRepository.CheckForUserExistence(userName, userPassword)
}

func (s *AuthService) GenerateToken(name string, password string) (string, error) {
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

func (s *AuthService) ParseToken(accessToken string) (int, error) {

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
