package db

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func gormConnect() *gorm.DB {
	DBMS := os.Getenv("SASG_DBMS")
	USER := os.Getenv("SASG_USER")
	PASS := os.Getenv("SASG_PASS")
	DBNAME := os.Getenv("SASG_DBNAME")
	// postfix 'parse...' for charcode of mysql
	CONNECT := USER + ":" + PASS + "@/" + DBNAME + "?parseTime=true"

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}

	return db
}

// Initialize DB
func Init() {
	db := gormConnect()

	defer db.Close()
	db.AutoMigrate(&Article{})
	db.AutoMigrate(&Tag{})
	db.AutoMigrate(&User{})
}
