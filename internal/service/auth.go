package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"marketplace/internal/models"
	"marketplace/internal/repository"
	"os"
	"time"
)

const (
	salt = "asklfjn2jdnalkmsd"
	//signingKey = "adSj23&h#!kjWjqwnd@jnef7832N"
	tokenTTL = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.MapClaims
	Login string `json:"login"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) CreateUser(user models.User) (models.User, error) {
	if err := user.Validate(); err != nil {
		return models.User{}, err
	}
	hashedPassword, err := generatePassword(user.Password)
	if err != nil {
		return models.User{}, err
	}
	user.Password = hashedPassword
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(login, password string) (string, error) {
	user, err := s.repo.GetUser(login)
	if err != nil {
		return "", err
	}
	if err := verifyPassword(user.Password, password); err != nil {
		return "", errors.New("неверный пароль")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.MapClaims{
			"exp": time.Now().Add(tokenTTL).Unix(),
			"iat": time.Now().Unix(),
		},
		user.Login,
	})

	return token.SignedString([]byte(getSigningKey()))
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Неверный метод авторизации")
		}
		return []byte(getSigningKey()), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("клеймы токена не типа  *tokenClaims")
	}

	return claims.Login, nil
}

func generatePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func verifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func getSigningKey() string {
	return os.Getenv("JWT_SIGNING_KEY")
}
