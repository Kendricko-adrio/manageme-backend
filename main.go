package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kendricko-adrio/to-do-backend/controller"
	"github.com/rs/cors"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hallo")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", hello)
	r.HandleFunc("/login", controller.Login).Methods(http.MethodPost)
	r.HandleFunc("/login", controller.GetLogin).Methods(http.MethodGet)
	r.HandleFunc("/task/to-do/{task_type_id}", controller.GetTaskToDo).Methods(http.MethodGet)
	r.HandleFunc("/task/to-do/{task_type_id}", controller.SetTaskToDo).Methods(http.MethodPost)
	r.HandleFunc("/task/to-do/{task_id}/right", controller.TaskShiftRight).Methods(http.MethodPut)
	r.HandleFunc("/task/to-do/{task_id}/left", controller.TaskShiftLeft).Methods(http.MethodPut)
	r.HandleFunc("/schedules", controller.GetAllUserSchedule).Methods(http.MethodGet)
	r.HandleFunc("/schedule/add", controller.AddSchedule).Methods(http.MethodPost)
	// r.HandleFunc("/user/{id}", controller.GetUserById).Methods(http.MethodGet)
	// migrate.Migrate()
	// r.Use(middleware.CorsMiddleware)

	c := cors.New(cors.Options{
		// AllowedOrigins: []string{"http://localhost:3000"},
		AllowedOrigins:   []string{"https://managemesite.netlify.app"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "PUT", "POST", "PATCH", "OPTION"},
		AllowedHeaders:   []string{"Content-Type", "Accept", "Authorization", "Content-Length", "Accept-Encoding"},
	})

	handler := cors.Default().Handler(r)

	// handler := c.Handler(r)
	// http.Handle("/", r)
	if err := http.ListenAndServe("localhost:1234", handler); err != nil {
		fmt.Print(err)
	}

}
