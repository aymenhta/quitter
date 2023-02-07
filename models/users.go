package models

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int
	Username       string
	Email          string
	HashedPassword []byte
	JoinedAt       time.Time
}

type UserModel struct {
	DB *pgxpool.Pool
}

func (model *UserModel) Create(username, email, password string) (int, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return 0, err
	}

	stmnt := `INSERT INTO users (username, email, password, joined_at) 
					VALUES($1, $2, $3, CURRENT_DATE)
					RETURNING id;`

	var id int
	err = model.DB.QueryRow(context.Background(), stmnt, username, email, hashedPassword).Scan(&id)
	if err != nil {
		errMsg := err.Error()
		if strings.Contains(errMsg, pgerrcode.UniqueViolation) {
			if strings.Contains(errMsg, "users_email_key") {
				return 0, errors.New("email already taken")
			} else if strings.Contains(errMsg, "users_username_key") {
				return 0, errors.New("username already taken")
			}
		}
		return 0, err
	}

	return id, nil
}
