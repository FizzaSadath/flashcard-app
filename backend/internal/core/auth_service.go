package core

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo  UserRepository
	jwtSecret []byte
}

func NewAuthService(repo UserRepository, secret string) *AuthService {
	return &AuthService{
		userRepo:  repo,
		jwtSecret: []byte(secret),
	}
}

func (s *AuthService) Register(email, username, password, confirmPassword string) error {

	existingEmail, _ := s.userRepo.GetUserByEmail(email)
	if existingEmail != nil {
		return errors.New("email already registered")
	}
	existingUser, _ := s.userRepo.GetUserByUsername(username)
	if existingUser != nil {
		return errors.New("username is already taken")
	}
	if password == confirmPassword {

		hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
		if err != nil {
			return err
		}

		user := &User{
			Email:    email,
			Username: username,
			Password: string(hashedBytes), // save hashed password
		}

		return s.userRepo.CreateUser(user)
	} else {
		return errors.New("passwords didn't match")
	}
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"email":    user.Email,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // Valid for 3 days
	})

	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
