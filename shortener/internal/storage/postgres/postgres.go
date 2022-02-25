package postgres

import (
	"context"
	"github.com/GinGin3203/shorturl/shortener/internal/service_errors"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type db struct {
	*pgx.Conn
}

func NewConn(ctx context.Context, url string) (*db, error) {
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		return nil, errors.Wrap(err, "NewConn")
	}
	stmtCreateTable := `CREATE TABLE IF NOT EXISTS url_table (
		id SERIAL PRIMARY KEY,
		long_url VARCHAR UNIQUE
	)`

	_, err = conn.Exec(ctx, stmtCreateTable)
	if err != nil {
		return nil, errors.Wrap(err, "NewConn")
	}
	return &db{conn}, nil
}

func (db *db) GetURLByID(ctx context.Context, id int) (string, error) {
	stmtQuery := "SELECT (long_url) FROM url_table WHERE id=$1"
	row := db.QueryRow(ctx, stmtQuery, id)
	var url string
	if err := row.Scan(&url); err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return "", service_errors.ErrShortURLNotFound
		default:
			return "", err
		}
	}
	return url, nil
}

func (db *db) InsertAndGetID(ctx context.Context, url string) (int, error) {
	// Можно было бы обойтись одним SQL-запросом и не ходить в базу 2 раза
	stmtCheckIfExists := "INSERT INTO url_table(long_url) VALUES ($1) ON CONFLICT DO NOTHING"
	_, err := db.Exec(ctx, stmtCheckIfExists, url)
	if err != nil {
		return 0, errors.Wrap(err, "InsertAndGetID")
	}
	stmtQuery := "SELECT (id) FROM url_table WHERE long_url=$1"
	row := db.QueryRow(ctx, stmtQuery, url)
	var id int
	if err = row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
