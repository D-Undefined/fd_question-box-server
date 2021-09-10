package repository

import "github.com/kazuki-komori/fd_question-box-server/domain/entity"

type QuestionRepository interface {
	CreateQuestion(question *entity.Question) error
	GetByID(int) (*entity.Question, error)
	FindAll(*[]entity.Question) error
}