package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"

	"gotasks/internal/repository/storage"
)

const (
	UsersTable string = "s_users"
	DataTable         = "data_usecase"
)

func Connect(dsn string) (storage.StorageModel, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return storage.StorageModel{}, err
	}

	if err = db.Ping(); err != nil {
		return storage.StorageModel{}, err
	}

	return storage.StorageModel{DB: db}, nil
}
