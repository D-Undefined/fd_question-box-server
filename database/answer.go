package database

import (
	"time"

	"github.com/kazuki-komori/fd_question-box-server/domain/entity"
	"github.com/kazuki-komori/fd_question-box-server/domain/repository"
)


type AnswerRepository struct {
	SqlHandler
}

func NewAnswerRepository(sh SqlHandler) repository.AnswerRepository {
	answerRepository := AnswerRepository{sh}
	return &answerRepository
}

func (aR AnswerRepository) Create(answer *entity.Answer) error {
	db := aR.SqlHandler.db
	answer.CreatedAt = time.Now().Unix()
	err := db.Create(answer).Error
	if err != nil {
		return err
	}
	question := new(entity.Question)
	question.ID = answer.QuestionID
	err = db.Model(question).Update("isAnswered", true).Error
	if err != nil {
		return err
	}
	return nil
}
