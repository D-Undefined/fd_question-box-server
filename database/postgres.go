package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type SqlHandler struct {
	db *gorm.DB
}


func NewDB() (*SqlHandler, error) {

	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		return nil, fmt.Errorf("failed to connect db= %w", err)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.db = db
	return sqlHandler, nil
}