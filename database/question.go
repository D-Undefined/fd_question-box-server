package database

import (
	"fmt"
	"time"

	"github.com/kazuki-komori/fd_question-box-server/domain/entity"
	"github.com/kazuki-komori/fd_question-box-server/domain/repository"
)

type QuetionRepository struct {
	SqlHandler
}

func NewQuestionRepository(sh SqlHandler) repository.QuestionRepository {
	questionRepository := QuetionRepository{sh}
	return &questionRepository
}

// 質問を作成
func (qR *QuetionRepository) CreateQuestion(question *entity.Question) (err error) {
	db := qR.SqlHandler.db
	question.CreatedAt = time.Now().Unix()
	return db.Create(&question).Error
}

// 質問を取得
func (qR *QuetionRepository) GetByID(id int) (*entity.Question, error) {
	db := qR.SqlHandler.db
	dto := new(entity.Question)
	dto.ID = id
	err := db.Preload("Answer").First(dto).Error
	if err != nil {
		fmt.Printf("%v", err)
		return nil, err
	}
	return dto, nil
}