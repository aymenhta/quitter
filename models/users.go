package models

import (
	"context"
	"errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *pgxpool.Pool
}

func (model *UserModel) Insert(username, email, password string) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmnt := `INSERT INTO users (username, email, password, created) 
			VALUES(?, ?, ?, CURRENT_DATE)`

	_, err = model.DB.Exec(context.Background(), stmnt, username, email, hashedPassword)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			// Here check for duplicate emails or userames code 23505
			if pgErr.Code == pgerrcode.UniqueViolation {
				return errors.New("email/username already exists")
			}
			// fmt.Println(pgErr.Message) // => syntax error at end of input
			// fmt.Println(pgErr.Code)    // => 42601
		}
		return err
	}
	return nil
}
