package userHandler

import (
	"fmt"
	db2 "github.com/hize8/login-api/db"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func saveNewUser(newUser User) {
	db, err := gorm.Open("postgres", db2.GetUrlConnection())
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	db.Create(&newUser)
}

func getUserByEmail(email string) User {
	db, err := gorm.Open("postgres", db2.GetUrlConnection())
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	var u = User{}
	db.Where("email = ?", email).First(&u)
	return u
}

func getAllUsers() []User {
	db, err := gorm.Open("postgres", db2.GetUrlConnection())
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	var users []User
	db.Find(&users)
	return users
}
