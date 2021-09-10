package repository

import "github.com/kazuki-komori/fd_question-box-server/domain/entity"


type AnswerRepository interface {
	Create(*entity.Answer) error
}