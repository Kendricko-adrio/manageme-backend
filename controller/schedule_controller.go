package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/kendricko-adrio/to-do-backend/database"
	"github.com/kendricko-adrio/to-do-backend/model"
)

func GetAllUserSchedule(w http.ResponseWriter, r *http.Request) {

	user, err := GetUserFromCookie(w, r)
	if err != nil {
		WriteResponse(w, r, http.StatusNoContent, err)
		return
	}
	db := database.GetInstance()

	schedules := []model.Schedule{}

	result := db.Where("user_id = ?", user.ID).Find(&schedules)

	if result.RowsAffected == 0 {
		WriteResponse(w, r, http.StatusNoContent, "No Data Found")
	}

	WriteResponse(w, r, http.StatusOK, schedules)
}

func AddSchedule(w http.ResponseWriter, r *http.Request) {

	user, err := GetUserFromCookie(w, r)
	if err != nil {
		panic(err)
	}

	desc := r.FormValue("description")
	deadline, err := strconv.ParseInt(r.FormValue("deadline"), 10, 64)

	if err != nil {
		panic(err)
	}

	time := time.UnixMilli(deadline)

	db := database.GetInstance()

	db.Create(&model.Schedule{
		Description:    desc,
		ScheduleTypeId: 1,
		Deadline:       time,
		UserId:         user.ID,
	})

	WriteResponse(w, r, 200, "success")
}
