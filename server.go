package main

import (
	"fmt"
	"github.com/hize8/login-api/db"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/hize8/login-api/userHandler"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	configureDb()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/sign-up", userHandler.SignUp)
	e.POST("/log-in", userHandler.Login)

	r := e.Group("/users")
	r.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	r.GET("/all", userHandler.GetUsers)

	e.Logger.Fatal(e.Start(":1323"))
}

func configureDb() {
	db, err := gorm.Open("postgres", db.GetUrlConnection())
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Aquí van las migraciones
	db.AutoMigrate(&userHandler.User{})
}
