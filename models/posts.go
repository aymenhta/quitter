package models

import (
	"context"
	"errors"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostsList struct {
	ID       int
	Content  string
	PostedAt time.Time
	UserId   int
	Username string
}

type PostDetails struct {
	ID       int
	Content  string
	PostedAt time.Time
	UserId   int
	Username string
	// Replies  []PostDetails
}

type PostModel struct {
	DB *pgxpool.Pool
}

func (model *PostModel) Create(content string, userId int) (int, time.Time, error) {
	stmnt := `INSERT INTO posts (content, user_id, posted_at)
					VALUES ($1, $2, NOW())
					RETURNING id, posted_at;`

	var id int
	var postedAt pgtype.Timestamp
	err := model.DB.QueryRow(context.Background(), stmnt, content, userId).Scan(&id, &postedAt)
	if err != nil {
		return 0, time.Time{}, err
	}

	return id, postedAt.Time, nil
}

func (model *PostModel) GetPosts() ([]*PostsList, error) {
	stmnt := `SELECT P.id, P.content, P.posted_at,
				P.user_id, U.username
				FROM posts P JOIN users U ON P.user_id = U.id
				ORDER BY P.posted_at DESC;`

	// row := model.DB.QueryRow(context.Background(), stmnt, id)

	var posts []*PostsList
	err := pgxscan.Select(context.Background(), model.DB, &posts, stmnt)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (model *PostModel) GetPost(id int) (*PostDetails, error) {
	stmnt := `SELECT P.id, P.content, P.posted_at,
				P.user_id, U.username
				FROM posts P JOIN users U ON P.user_id = U.id
				WHERE P.id = $1;`

	// row := model.DB.QueryRow(context.Background(), stmnt, id)

	postDetails := &PostDetails{}
	err := pgxscan.Get(context.Background(), model.DB, postDetails, stmnt, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}

	return postDetails, nil
}

func (model *PostModel) DeletePost(id int) error {
	stmnt := `DELETE FROM posts WHERE id = $1;`

	commandTag, err := model.DB.Exec(context.Background(), stmnt, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return ErrNoRecord
	}

	return nil
}

func (model *PostModel) PostExists(id int) (bool, error) {
	var exists bool
	stmt := "SELECT EXISTS(SELECT true FROM posts WHERE id = $1)"
	err := model.DB.QueryRow(context.Background(), stmt, id).Scan(&exists)
	return exists, err
}

func (model *PostModel) GetUserPosts(userId int) ([]*PostsList, error) {
	stmnt := `SELECT P.id, P.content, P.posted_at,
				P.user_id, U.username
				FROM posts P JOIN users U ON P.user_id = U.id
				WHERE P.user_id = $1
				ORDER BY P.posted_at DESC;`

	// row := model.DB.QueryRow(context.Background(), stmnt, id)

	var posts []*PostsList
	err := pgxscan.Select(context.Background(), model.DB, &posts, stmnt, userId)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
