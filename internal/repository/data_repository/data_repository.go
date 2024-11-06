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

func NewDataRepository(storage storage.StorageModel) *DataRepository {
	return &DataRepository{
		Storage: storage,
	}
}

func (dr *DataRepository) AddData(
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

func (ur *DataRepository) GetRandomData(
	ctx context.Context,
	uid uint64,
) (d entity.Data, err error) {
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

	rows, err := ur.Storage.DB.Query(q)
	if err != nil {
		return
	}
	defer rows.Close()

	data := entity.Data{}

	for rows.Next() {
		if err = rows.Scan(
			&data.ID,
			&data.User,
			&data.Link,
			&data.Note,
			&data.AddedAt,
		); err != nil {
			return entity.Data{}, err
		}
	}

	return data, nil
}

func (ur *DataRepository) GetAllData() (all_data []entity.Data, err error) {
	q := fmt.Sprintf("SELECT * FROM %s", postgres.DataTable)

	rows, err := ur.Storage.DB.Query(q)
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

func (ur *DataRepository) DeleteData(
	ctx context.Context,
	did uint64,
) {

}
