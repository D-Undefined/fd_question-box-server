package entity


type Answer struct {
	ID int `json:"id" gorm:"primary_key"`
	Content  string `json:"content"`
	QuestionID  int `json:"question_id"`
	CreatedAt int64 `json:"created_at"`
}