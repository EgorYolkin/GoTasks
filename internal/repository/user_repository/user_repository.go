package userrepository

import (
	"context"
	"fmt"

	"gotasks/internal/entity"
	"gotasks/internal/repository/postgres"
	"gotasks/internal/repository/storage"
)

type UserRepository struct {
	Storage storage.StorageModel
}

func NewUserRepository(storage storage.StorageModel) *UserRepository {
	return &UserRepository{
		Storage: storage,
	}
}

func (ur *UserRepository) CreateUser(
	ctx context.Context,
	u *entity.User,
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
		u.TelegramId, u.CreatedAt,
	)

	if _, err := ur.Storage.DB.Exec(q); err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) GetUser(
	ctx context.Context,
	uid uint64,
) {

}

func (ur *UserRepository) GetAllUser() (users []entity.User, err error) {
	q := fmt.Sprintf("SELECT * FROM %s", postgres.UsersTable)

	rows, err := ur.Storage.DB.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := entity.User{}
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

func (ur *UserRepository) DeleteUser(
	ctx context.Context,
	uid uint64,
) {

}
