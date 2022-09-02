package main

import (
	"backend/database"
	m "backend/models"
	"backend/routes"
	"fmt"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

func initDatabase() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		"root",
		"",
		"127.0.0.1",
		"3306",
		"golang_test",
	)
	var err error
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected!")
	database.DBConn.AutoMigrate(&m.Dogs{})
	database.DBConn.AutoMigrate(&m.Register{})
	database.DBConn.AutoMigrate(&m.Profile_User{})
}

func main() {
	app := fiber.New()
	initDatabase()
	routes.UserRoute(app)

}

// From github.com dsfsdf agin/
