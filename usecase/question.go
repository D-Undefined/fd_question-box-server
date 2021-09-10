package usecase

import (
	"fmt"

	"github.com/kazuki-komori/fd_question-box-server/domain/entity"
	"github.com/kazuki-komori/fd_question-box-server/domain/repository"
)

type QuestionUC interface {
	Create(*entity.Question) (*entity.Question, error)
	GetByID(int) (*entity.Question, error)
	FindAll() (*[]entity.Question, error)
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

func (uc *questionUC) GetByID(id int) (*entity.Question, error) {
	return uc.questionRepo.GetByID(id)
}

func (uc *questionUC) FindAll() (*[]entity.Question, error) {
	questions := new([]entity.Question)
	err := uc.questionRepo.FindAll(questions)
	if err != nil {
		return nil, fmt.Errorf("questionRepo CreateQuestion: %v", err)
	}
	return questions, nil
}

