package userHandler

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" query:"name"`
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
}
