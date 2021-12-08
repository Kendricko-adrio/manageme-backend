package model

import (
	"github.com/kendricko-adrio/to-do-backend/database"
	"gorm.io/gorm"
)

type TaskType struct {
	gorm.Model
	TaskTypeName string `json:"task_type_name"`
	Task         []Task `json:"task" gorm:"foreignKey:TaskTypeId"`
}

func SeedTaskType() {
	db := database.GetInstance()
	db.Create(&TaskType{
		TaskTypeName: "To Do",
	})
	db.Create(&TaskType{
		TaskTypeName: "On Progress",
	})
	db.Create(&TaskType{
		TaskTypeName: "Completed",
	})
}
