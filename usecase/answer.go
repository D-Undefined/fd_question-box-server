package usecase

import (
	"github.com/kazuki-komori/fd_question-box-server/domain/entity"
	"github.com/kazuki-komori/fd_question-box-server/domain/repository"
)


type AnswerUC interface {
	Create(*entity.Answer)(*entity.Answer, error)
}

type answerUC struct {
	answerRepo repository.AnswerRepository
}

func NewAnswerUC(aR repository.AnswerRepository) AnswerUC {
	answerUC := answerUC{answerRepo: aR}
	return &answerUC
}

func (aU answerUC) Create(answer *entity.Answer)(*entity.Answer, error) {
	err := aU.answerRepo.Create(answer)
	return answer, err
}

