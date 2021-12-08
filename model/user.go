package model

import (
	"github.com/kendricko-adrio/to-do-backend/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string     `json:"username"`
	Password string     `json:"password"`
	Task     []Task     `json:"task" gorm:"foreignKey:UserId"`
	Schedule []Schedule `json:"schedule" gorm:"foreignKey:UserId"`
}

func SeedUser() {
	db := database.GetInstance()
	db.Create(&User{
		Username: "ricko",
		Password: "pacinto",
	})

}
