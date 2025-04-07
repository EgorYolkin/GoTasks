package user_repository

import (
	"context"
	"fmt"
	"gotasks/internal/repository/database/postgres"

	"gotasks/internal/domain"
	"gotasks/internal/repository/storage"
)

type PostgresUserRepository struct {
	Storage storage.StorageModel
}

func NewRepository(storage storage.StorageModel) *PostgresUserRepository {
	return &PostgresUserRepository{
		Storage: storage,
	}
}

func (ur *PostgresUserRepository) Create(
	ctx context.Context,
	user domain.User,
) error {
	q := fmt.Sprintf(`
    	INSERT INTO %s
        (telegram_id, created_at)
        VALUES
        (%d, %d)
    	ON CONFLICT (telegram_id)
    	    DO NOTHING;
    	`,
		postgres.UsersTable,
		user.TelegramId, user.CreatedAt,
	)

	if _, err := ur.Storage.DB.Exec(q); err != nil {
		return err
	}
	return nil
}

func (ur *PostgresUserRepository) GetAll(ctx context.Context) (users []domain.User, err error) {
	q := fmt.Sprintf("SELECT * FROM %s", postgres.UsersTable)

	rows, err := ur.Storage.DB.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := domain.User{}
		if err = rows.Scan(
			&user.ID,
			&user.TelegramId,
		); err != nil {
			return
		}
		users = append(users, user)
	}
	return
}
