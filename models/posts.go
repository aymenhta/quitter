package models

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Post struct {
	ID       int
	Content  string
	PostedAt time.Time
	UserId   int
}

type PostModel struct {
	DB *pgxpool.Pool
}

func (model *PostModel) Insert(content string, userId int) (int, string, error) {
	stmnt := `INSERT INTO posts (content, user_id, posted_at)
					VALUES ($1, $2, NOW())
					RETURNING id, posted_at;`

	var id int
	var postedAt pgtype.Timestamp
	err := model.DB.QueryRow(context.Background(), stmnt, content, userId).Scan(&id, &postedAt)
	if err != nil {
		return 0, "", err
	}

	humanizedTimestamp := ToHumanTimeStamp(postedAt.Time)
	fmt.Printf("created at: %v\n", humanizedTimestamp)
	return id, humanizedTimestamp, nil
}

func ToHumanTimeStamp(timestamp time.Time) string {
	return timestamp.UTC().Format("02 Jan 2006 at 15:04")
}
