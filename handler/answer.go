package handler

import (
	"fmt"
	"net/http"

	"github.com/kazuki-komori/fd_question-box-server/domain/entity"
	"github.com/kazuki-komori/fd_question-box-server/usecase"
	"github.com/labstack/echo/v4"
)



type AnswerHandler interface {
	PutAnswer(echo.Context) error
}

type answerHandler struct {
	answerUC usecase.AnswerUC
}

func NewAnswerHandler(aU usecase.AnswerUC) AnswerHandler {
	return &answerHandler{answerUC: aU}
}

func (aH answerHandler) PutAnswer(c echo.Context) error {
	dto := new(entity.Answer)
	c.Bind(dto)
	e, err := aH.answerUC.Create(dto)
	if err != nil {
		fmt.Printf("%v", err)
		return c.JSON(http.StatusInternalServerError, APIError{Message: "パラメータが不正です"})
	}
	return c.JSON(http.StatusCreated, e)
}

