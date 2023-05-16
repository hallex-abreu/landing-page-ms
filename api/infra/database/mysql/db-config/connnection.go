package dbconfig

import (
	"fmt"
	"os"

	"github.com/hallex-abreu/landing-page-ms/api/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Mysql *gorm.DB

func Connection() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database!")
	}

	db.AutoMigrate(&domain.Lead{})

	Mysql = db
}
