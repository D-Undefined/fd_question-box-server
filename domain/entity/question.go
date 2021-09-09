package entity

type Question struct {
	ID uint32 `json:"id"`
	Content  string `json:"content"`
	IsAnswerd  bool `json:"is_answerd"`
	CreatedAt int64 `json:"created_at"`
}