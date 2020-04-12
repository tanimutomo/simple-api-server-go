package db

import (
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type ErrorResponse struct {
	IsError bool
	Status  int
	Message string
}

func gormConnect() (*gorm.DB, ErrorResponse) {
	DBMS := os.Getenv("SASG_DBMS")
	USER := os.Getenv("SASG_USER")
	PASS := os.Getenv("SASG_PASS")
	DBNAME := os.Getenv("SASG_DBNAME")
	// postfix 'parse...' for charcode of mysql
	CONNECT := USER + ":" + PASS + "@/" + DBNAME + "?parseTime=true"

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		return nil,
			ErrorResponse{
				IsError: true,
				Status:  http.StatusInternalServerError,
				Message: "Cannot access to database. " + err.Error(),
			}
	}
	return db, ErrorResponse{IsError: false}
}

// Initialize DB
func Init() ErrorResponse {
	db, errResp := gormConnect()
	if errResp.IsError {
		return errResp
	}

	defer db.Close()
	db.AutoMigrate(&Article{})
	db.AutoMigrate(&Tag{})
	db.AutoMigrate(&User{})

	return ErrorResponse{IsError: false}
}
