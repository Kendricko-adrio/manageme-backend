package model

import (
	"github.com/kendricko-adrio/to-do-backend/database"
	"gorm.io/gorm"
)

type ScheduleType struct {
	gorm.Model
	TypeDescription string
	Schedule        []Schedule `gorm:"foreignKey:ScheduleTypeId"`
}

func ScheduleTypeSeed() {
	db := database.GetInstance()
	db.Create(&ScheduleType{
		TypeDescription: "Not Done",
	})
	db.Create(&ScheduleType{
		TypeDescription: "Done",
	})
}
