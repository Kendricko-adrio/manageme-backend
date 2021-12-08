package controller

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, r *http.Request, code int, content interface{}) {

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(content); err != nil {
		panic(err)
	}
}
