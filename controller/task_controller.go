package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kendricko-adrio/to-do-backend/database"
	"github.com/kendricko-adrio/to-do-backend/model"
	"gorm.io/gorm"
)

func TaskShiftRight(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(r)["task_id"])
	db := database.GetInstance()
	task := &model.Task{
		Model: gorm.Model{
			ID: uint(id),
		},
	}

	db.Find(task)

	if task.TaskTypeId == 3 {
		task.TaskTypeId = 1
	} else {
		task.TaskTypeId = task.TaskTypeId + 1
	}
	db.Save(task)

	WriteResponse(w, r, http.StatusOK, task)
	// fmt.Println(task)
}

func TaskShiftLeft(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(r)["task_id"])
	db := database.GetInstance()
	task := &model.Task{
		Model: gorm.Model{
			ID: uint(id),
		},
	}

	db.Find(task)

	if task.TaskTypeId == 1 {
		task.TaskTypeId = 3
	} else {
		task.TaskTypeId = task.TaskTypeId - 1
	}
	db.Save(task)

	WriteResponse(w, r, http.StatusOK, task)
	// fmt.Println(task)
}

func GetTaskToDo(w http.ResponseWriter, r *http.Request) {

	token, err := r.Cookie("manageme")

	// fmt.Printf("token: %v", token)

	if err != nil {
		WriteResponse(w, r, http.StatusNotFound, err)
		return
	}

	user, err := GetUserByToken(token.Value)

	if err != nil {
		WriteResponse(w, r, http.StatusNotFound, err)
	}

	vars := mux.Vars(r)
	task_type_id := vars["task_type_id"]
	db := database.GetInstance()
	tasks := []model.Task{}
	db.Where("task_type_id = ? AND user_id = ? ", task_type_id, user.ID).Find(&tasks)
	// fmt.Printf("task : %v", tasks)
	WriteResponse(w, r, http.StatusOK, &tasks)
}

func SetTaskToDo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := r.FormValue("title")
	description := r.FormValue("description")
	task_type_id, _ := strconv.Atoi(vars["task_type_id"])
	db := database.GetInstance()

	task := &model.Task{
		TaskTitle:       title,
		TaskDescription: description,
		TaskTypeId:      uint(task_type_id),
		UserId:          1,
	}
	db.Create(task)
	// fmt.Println(task)
	// fmt.FPrintf(task)

	tasks := []model.Task{}
	db.Where("task_type_id = ?", task_type_id).Find(&tasks)
	// fmt.Println(tasks)
	WriteResponse(w, r, http.StatusOK, &tasks)
}
