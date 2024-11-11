package data_repository

import (
	"context"
	"fmt"

	"gotasks/internal/entity"
	"gotasks/internal/repository/postgres"
	"gotasks/internal/repository/storage"
)

type DataRepository struct {
	Storage storage.StorageModel
}

func NewRepository(storage storage.StorageModel) *DataRepository {
	return &DataRepository{
		Storage: storage,
	}
}

func (dr *DataRepository) Create(
	ctx context.Context,
	d entity.Data,
) error {
	q := fmt.Sprintf(`
    	INSERT INTO %s
        ("user", link, note, added_at)
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

func (dr *DataRepository) GetOne(
	ctx context.Context,
	uid uint64,
) (d *entity.Data, err error) {
	q := fmt.Sprintf(
		`
            SELECT *
            FROM
                %s
            WHERE
                "user"=%d
            ORDER BY RANDOM()
            LIMIT 1;
        `,
		postgres.DataTable,
		int(uid),
	)

	rows, err := dr.Storage.DB.Query(q)
	if err != nil {
		return
	}
	defer rows.Close()

	data := entity.Data{}

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
		return nil, fmt.Errorf("Data not found")
	}

	return &data, nil
}

func (dr *DataRepository) GetAll() (all_data []entity.Data, err error) {
	q := fmt.Sprintf("SELECT * FROM %s", postgres.DataTable)

	rows, err := dr.Storage.DB.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		data := entity.Data{}
		if err = rows.Scan(
			&data.ID,
			&data.User,
			&data.Link,
			&data.Note,
			&data.AddedAt,
		); err != nil {
			return
		}
		all_data = append(all_data, data)
	}
	return
}

func (dr *DataRepository) Delete(
	ctx context.Context,
	did uint64,
) error {
	q := fmt.Sprintf(`
       	DELETE FROM %s
        WHERE id=%d
       	`,
		postgres.DataTable,
		did,
	)

	if _, err := dr.Storage.DB.Exec(q); err != nil {
		return err
	}
	return nil
}
