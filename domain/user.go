package domain

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	UserID int    `json:"userID"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func (u *User) Validate() error {
	if err := validateUsername(u.Username); err != nil {
		return err
	}
	if err := validatePassword(u.Password); err != nil {
		return err
	}
	if err := validateRole(u.Role); err != nil {
		return err
	}
	return nil
}

func validateUsername(username string) error {
	if strings.TrimSpace(username) == "" {
		return errors.New("username must not be empty")
	}
	if len(username) < 3 {
		return errors.New("username must be at least 3 characters long")
	}
	if len(username) > 50 {
		return errors.New("username must be less than 50 characters long")
	}
	return nil
}

func (c *Credentials) Validate() error {
	if err := validateUsername(c.Username); err != nil {
		return err
	}
	if err := validatePassword(c.Password); err != nil {
		return err
	}
	return nil
}

func validatePassword(password string) error {
	if strings.TrimSpace(password) == "" {
		return errors.New("password must not be empty")
	}
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}
	return nil
}

func validateRole(role string) error {
	if role != "user" && role != "admin" {
		return errors.New("role must be either 'user' or 'admin'")
	}
	return nil
}

// GenerateJWT generates a new JWT token
func GenerateJWT(user User) (string, error) {
	claims := &Claims{
		UserID: user.ID,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("SECRET_KEY"))
}

// ValidateJWT validates a JWT token and returns the claims
func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("SECRET_KEY"), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

type UserRepository interface {
	CreateUser(user *User) error
	AuthenticateUser(username, password string) (string, error)
	GetUserByUsername(username string) (*User, error)
	ValidateToken(tokenString string) (*Claims, error)
}
