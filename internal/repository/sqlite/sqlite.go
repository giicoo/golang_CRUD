package repository

import (
	"database/sql"
	"os"

	"github.com/giicoo/golang_CRUD/internal/config"
)

type Sqlite struct {
	db  *sql.DB
	cfg *config.Config
}

func NewRepo(db *sql.DB, cfg *config.Config) *Sqlite {
	return &Sqlite{
		db:  db,
		cfg: cfg,
	}
}

func (sqlite *Sqlite) InitDB() error {
	path := sqlite.cfg.Path_to_sql + "/create_books.sql"
	create_book_stmt, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	path = sqlite.cfg.Path_to_sql + "/create_authors.sql"
	create_author_stmt, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	stmt := string(create_book_stmt)
	_, err = sqlite.db.Exec(stmt)
	if err != nil {
		return err
	}

	stmt = string(create_author_stmt)
	_, err = sqlite.db.Exec(stmt)
	if err != nil {
		return err
	}

	return sqlite.db.Ping()
}
