package repository

import (
	"os"

	"github.com/giicoo/golang_CRUD/internal/models"
)

func (sqlite *Sqlite) AddAuthor(name string) (int, error) {

	path := sqlite.cfg.Path_to_sql + "/insert/insert_author.sql"

	stmt, err := os.ReadFile(path)
	if err != nil {
		return -1, err
	}

	res, err := sqlite.db.Exec(string(stmt), name)
	if err != nil {
		return -1, err
	}

	author_id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(author_id), nil
}

func (sqlite *Sqlite) GetAuthor(id int, name string) ([]models.Author, error) {

	path := sqlite.cfg.Path_to_sql + "/select/select_author.sql"

	stmt, err := os.ReadFile(path)
	author := models.Author{}

	if err != nil {
		return nil, err
	}

	result := []models.Author{}

	rows, err := sqlite.db.Query(string(stmt), id, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&author.Author_id, &author.Name)
		if err != nil {
			return nil, err
		}
		result = append(result, author)
	}

	return result, nil
}

func (sqlite *Sqlite) DeleteAuthor(id int) error {
	path := sqlite.cfg.Path_to_sql + "/delete/delete_author.sql"

	stmt, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	_, err = sqlite.db.Exec(string(stmt), id)
	return err
}
