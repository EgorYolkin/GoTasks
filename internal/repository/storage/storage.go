package storage

import "database/sql"

type StorageModel struct {
	DB *sql.DB
}
