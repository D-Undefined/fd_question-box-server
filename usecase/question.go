package usecase

import (
	"fmt"

	"github.com/kazuki-komori/fd_question-box-server/domain/entity"
	"github.com/kazuki-komori/fd_question-box-server/domain/repository"
)

type QuestionUC interface {
	Create(*entity.Question) (*entity.Question, error)
}

type questionUC struct {
	questionRepo repository.QuestionRepository
}

func NewQuestionUC(qR repository.QuestionRepository) QuestionUC {
	questionUC := questionUC{questionRepo: qR}
	return &questionUC
}

func (uc *questionUC) Create(question *entity.Question) (*entity.Question, error) {
	err := uc.questionRepo.CreateQuestion(question)
	if err != nil {
		return nil, fmt.Errorf("questionRepo CreateQuestion: %v", err)
	}
	return question, nil
}