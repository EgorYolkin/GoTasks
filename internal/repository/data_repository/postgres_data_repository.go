package data_repository

import (
	"context"
	"fmt"
	"gotasks/internal/repository/database/postgres"

	"gotasks/internal/domain"
	"gotasks/internal/repository/storage"
)

type PostgresDataRepository struct {
	Storage storage.StorageModel
}

func NewPostgresDataRepository(storage storage.StorageModel) *PostgresDataRepository {
	return &PostgresDataRepository{
		Storage: storage,
	}
}

func (dr *PostgresDataRepository) Create(
	ctx context.Context,
	d domain.Data,
) error {
	q := fmt.Sprintf(`
    	INSERT INTO %s
        ("user_usecase", link, note, added_at)
        VALUES
        (%d, '%s', '%s', %d)
    	`,
		postgres.DataTable,
		d.User, d.Link, d.Note, d.AddedAt,
	)

	if _, err := dr.Storage.DB.Exec(q); err != nil {
		return err
	}
	return nil
}

func (dr *PostgresDataRepository) GetBy(
	ctx context.Context,
	key string,
	value int,
) (*domain.Data, error) {
	q := fmt.Sprintf(
		`
            SELECT *
            FROM
                %s
            WHERE
                %s=$1
            ORDER BY RANDOM()
            LIMIT 1;
        `,
		postgres.DataTable,
		key,
	)

	rows, err := dr.Storage.DB.Query(q, value)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	data := domain.Data{}

	if rows.Next() {
		if err = rows.Scan(
			&data.ID,
			&data.User,
			&data.Link,
			&data.Note,
			&data.AddedAt,
		); err != nil {
			return nil, err
		}
	} else {
		return nil, ErrDataNotFound
	}

	return &data, nil
}

func (dr *PostgresDataRepository) GetAll(ctx context.Context) ([]domain.Data, error) {
	q := fmt.Sprintf("SELECT * FROM %s", postgres.DataTable)

	rows, err := dr.Storage.DB.Query(q)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	data := []domain.Data{}

	for rows.Next() {
		dataObj := domain.Data{}
		if err = rows.Scan(
			&dataObj.ID,
			&dataObj.User,
			&dataObj.Link,
			&dataObj.Note,
			&dataObj.AddedAt,
		); err != nil {
			return nil, err
		}
		data = append(data, dataObj)
	}
	return data, nil
}

func (dr *PostgresDataRepository) DeleteBy(
	ctx context.Context,
	key string,
	value interface{},
) error {
	q := fmt.Sprintf(`
       	DELETE FROM %s
        WHERE %s=$1
       	`,
		postgres.DataTable,
		key,
	)

	if _, err := dr.Storage.DB.Exec(q, value); err != nil {
		return err
	}
	return nil
}
