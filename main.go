package main

import (
	"net/http"

	"github.com/kazuki-komori/fd_question-box-server/database"
	"github.com/kazuki-komori/fd_question-box-server/handler"
	"github.com/kazuki-komori/fd_question-box-server/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	db, _ := database.NewDB()
	qR := database.NewQuestionRepository(*db)
	aR := database.NewAnswerRepository(*db)
	
	qU := usecase.NewQuestionUC(qR)
	aU := usecase.NewAnswerUC(aR)

	qH := handler.NewQuestionHandler(qU)
	aH := handler.NewAnswerHandler(aU)

	// Echo instance
	e := echo.New()
  
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
  
	// Routes
	e.GET("/health", health)

	v1 := e.Group("/api/v1")

	v1.POST("/questions", qH.PostQuestion)
	v1.GET("/questions/:id", qH.GetQuestionByID)
	v1.GET("/questions", qH.GetQuestionAll)
	
	v1.PUT("/answer", aH.PutAnswer)
  
	// Start server
	e.Logger.Fatal(e.Start(":8080"))
  }
  
  // Handler
  func health(c echo.Context) error {
	return c.String(http.StatusOK, "Health OK ")
  }