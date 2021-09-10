package entity

type Question struct {
	ID int `json:"id" gorm:"primary_key"`
	Content  string `json:"content"`
	IsAnswerd  bool `json:"is_answerd"`
	CreatedAt int64 `json:"created_at"`
	Answer Answer `gorm:"foreignKey:QuestionID"`
}

type QuestionDTO struct {
	ID int `json:"id"`
	Content  string `json:"content"`
	IsAnswerd  bool `json:"is_answerd"`
	CreatedAt int64 `json:"created_at"`
}
