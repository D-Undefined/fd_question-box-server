package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/kazuki-komori/fd_question-box-server/domain/entity"
	"github.com/kazuki-komori/fd_question-box-server/usecase"
	"github.com/labstack/echo/v4"
)

type questionHandler struct {
	questionUC usecase.QuestionUC
}

type QuestionHandler interface {
	PostQuestion(c echo.Context) error
	GetQuestionByID(c echo.Context) error
	GetQuestionAll(c echo.Context) error
}

func NewQuestionHandler(qU usecase.QuestionUC) QuestionHandler {
	return &questionHandler{questionUC: qU}
}

// POST /questions 質問を作成
func (qH *questionHandler) PostQuestion (c echo.Context) error {
	question := new(entity.Question)
	c.Bind(question)
	e, err := qH.questionUC.Create(question)
	dto := entity.QuestionDTO{
		ID: e.ID,
		Content: e.Content,
		IsAnswerd: e.IsAnswerd,
		CreatedAt: e.CreatedAt,
	}
	if err != nil {
		fmt.Printf("%v", err)
		return c.JSON(http.StatusBadRequest, APIError{Message: "パラメータが不正です"})
	}
	return c.JSON(http.StatusOK, dto)
}

// GET /questions/:id 質問をIDで取得
func (qH *questionHandler)GetQuestionByID(c echo.Context) (err error) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, APIError{Message: "パラメータが不正です"})
	}
	e, err := qH.questionUC.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, APIError{Message: "サーバーでエラーが発生しました"})
	}
	if e.IsAnswerd == false {
		dto := entity.QuestionDTO{
			ID: e.ID,
			Content: e.Content,
			IsAnswerd: e.IsAnswerd,
			CreatedAt: e.CreatedAt,
		}
		return c.JSON(http.StatusOK, dto)
	}
	return c.JSON(http.StatusOK, e)
}

// GET /questions 質問を全件取得

func (qH *questionHandler)GetQuestionAll(c echo.Context) (err error) {
	e, err := qH.questionUC.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, APIError{Message: "サーバーでエラーが発生しました"})
	}
	return c.JSON(http.StatusOK, e)
}