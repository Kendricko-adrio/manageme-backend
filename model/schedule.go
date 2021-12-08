package model

import (
	"time"

	"github.com/kendricko-adrio/to-do-backend/database"
	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	Description    string        `json:"description"`
	Deadline       time.Time     `json:"deadline"`
	ScheduleTypeId uint          `json:"schedule_type_id"`
	ScheduleType   *ScheduleType `json:"schedule_type" gorm:"foreignKey:ScheduleTypeId"`
	UserId         uint          `json:"user_id"`
	User           *User         `gorm:"foreignKey:UserId"`
}

func ScheduleSeed() {

	db := database.GetInstance()
	db.Create(&Schedule{
		UserId:         1,
		ScheduleTypeId: 1,
		Description:    "Post makanan bung",
		Deadline:       time.Now().Add(time.Hour * 24),
	})

}
