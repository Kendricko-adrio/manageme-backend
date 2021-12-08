package migrate

import (
	"github.com/kendricko-adrio/to-do-backend/database"
	"github.com/kendricko-adrio/to-do-backend/model"
)

func Migrate() {
	db := database.GetInstance()
	db.Migrator().DropTable(&model.User{}, &model.TaskType{}, &model.Task{}, &model.ScheduleType{}, &model.Schedule{})
	db.AutoMigrate(&model.User{}, &model.TaskType{}, &model.Task{}, &model.ScheduleType{}, &model.Schedule{})
	seeder()
}

func seeder() {
	model.SeedUser()
	model.SeedTaskType()
	model.SeedTask()
	model.ScheduleTypeSeed()
	model.ScheduleSeed()
}
