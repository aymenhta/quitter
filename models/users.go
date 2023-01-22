package models

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgconn"
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

func (model *UserModel) Insert(username, email, password string) (int, error) {
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
		// TODO: FIX ERROR HANDLING
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) { // ? Why isn't it working...
			// Here check for duplicate emails or userames
			if pgErr.Code == pgerrcode.UniqueViolation {
				switch pgErr.ConstraintName {
				case "users_email_key":
					return 0, errors.New("email already taken")
				case "users_username_key":
					return 0, errors.New("username already taken")
				}
			}

			// fmt.Println(pgErr.Message) // => syntax error at end of input
			// fmt.Println(pgErr.Code)    // => 42601
		}
		return 0, err
	}

	return id, nil
}
