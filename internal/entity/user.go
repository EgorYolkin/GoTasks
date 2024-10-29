package entity

type User struct {
	ID         uint64 `json:"id"`
	TelegramId int    `json:"telegram_id"`
	CreatedAt  int    `json:"created_at"`
}
