package model

import (
	"github.com/kendricko-adrio/to-do-backend/database"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	TaskTitle       string    `json:"task_title"`
	TaskDescription string    `json:"task_description"`
	TaskTypeId      uint      `json:"task_type_id"`
	TaskType        *TaskType `json:"task_type"`
	UserId          uint      `json:"user_id"`
	User            *User     `json:"user" gorm:"foreignKey:UserId"`
}

func SeedTask() {
	db := database.GetInstance()
	db.Create(&Task{
		TaskTitle:       "Test",
		TaskDescription: "Tester aja",
		TaskTypeId:      1,
		UserId:          1,
	})
	db.Create(&Task{
		TaskTitle:       "Test",
		TaskDescription: "Tester aja",
		TaskTypeId:      2,
		UserId:          1,
	})
}
