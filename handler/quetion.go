package handler

import (
	"fmt"
	"net/http"

	"github.com/kazuki-komori/fd_question-box-server/domain/entity"
	"github.com/kazuki-komori/fd_question-box-server/usecase"
	"github.com/labstack/echo/v4"
)

type questionHandler struct {
	questionUC usecase.QuestionUC
}

type QuestionHandler interface {
	PostQuestion(c echo.Context) error
}

func NewQuestionHandler(qU usecase.QuestionUC) QuestionHandler {
	return &questionHandler{questionUC: qU}
}

func (qH *questionHandler) PostQuestion (c echo.Context) error {
	question := new(entity.Question)
	c.Bind(question)
	question, err := qH.questionUC.Create(question)
	if err != nil {
		fmt.Printf("%v", err)
		return c.JSON(http.StatusBadRequest, APIError{Message: "パラメータが不正です"})
	}
	return c.JSON(http.StatusOK, question)
}
