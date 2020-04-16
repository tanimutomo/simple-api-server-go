package db

import (
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type ErrorResponse struct {
	Status  int
	Message string
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

func gormConnect() (*gorm.DB, error) {
	DBMS := os.Getenv("SASG_DBMS")
	USER := os.Getenv("SASG_USER")
	PASS := os.Getenv("SASG_PASS")
	DBNAME := os.Getenv("SASG_DBNAME")
	// postfix 'parse...' for charcode of mysql
	CONNECT := USER + ":" + PASS + "@/" + DBNAME + "?parseTime=true"

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		return nil,
			&ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: "Cannot access to database. " + err.Error(),
			}
	}
	return db, nil
}

// Initialize DB
func Init() error {
	db, err := gormConnect()
	if err != nil {
		return err
	}

	defer db.Close()
	db.AutoMigrate(&Article{})
	db.AutoMigrate(&Tag{})
	db.AutoMigrate(&User{})

	return nil
}
