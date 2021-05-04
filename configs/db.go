package configs

import (
	"fmt"

	"crud_fiber_go_gorm/database"
	"crud_fiber_go_gorm/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() {

	var err error
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:password@tcp(localhost:3306)/local?charset=utf8mb4&parseTime=True&loc=Local"
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")

	defer database.DBConn.AutoMigrate(&models.Product{})

}
