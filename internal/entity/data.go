package entity

type Data struct {
	ID      uint64 `json:"id"`
	User    uint64 `json:"user_id"`
	Link    string `json:"link"`
	Note    string `json:"note"`
	AddedAt int    `json:"added_at"`
}
