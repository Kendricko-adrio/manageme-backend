package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kendricko-adrio/to-do-backend/database"
	"github.com/kendricko-adrio/to-do-backend/jwt"
	"github.com/kendricko-adrio/to-do-backend/model"
)

func Login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	db := database.GetInstance()

	user := &model.User{}

	result := db.Where("username like ? AND password like ?", username, password).Find(user)
	// fmt.Println(user)

	if result.RowsAffected == 0 {
		WriteResponse(w, r, http.StatusNotFound, "data not found")
		return
	}

	check, _ := r.Cookie("manageme")
	fmt.Println(check)

	token, _ := jwt.GenerateToken(user.ID)

	cookie := &http.Cookie{
		Name:     "manageme",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	}

	http.SetCookie(w, cookie)
	WriteResponse(w, r, http.StatusOK, user)
}

func GetUserFromCookie(w http.ResponseWriter, r *http.Request) (*model.User, error) {
	token, err := r.Cookie("manageme")

	if err != nil {
		return nil, err
	}

	user, err := GetUserByToken(token.Value)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetLogin(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("manageme")

	fmt.Println(token)

	if err != nil {
		WriteResponse(w, r, http.StatusNotFound, err)
		return
	}

	user, err := GetUserByToken(token.Value)

	if err != nil {
		WriteResponse(w, r, http.StatusNotFound, err)
	}

	WriteResponse(w, r, http.StatusOK, user)

}

func GetUserByToken(token string) (model.User, error) {

	if _, err := jwt.ParseToken(token); err != nil {
		return model.User{}, err
	}

	idToken, _ := jwt.ParseToken(token)
	db := database.GetInstance()

	user := &model.User{}

	result := db.Where("id = ?", idToken).Find(user)

	if result.RowsAffected == 0 {
		return model.User{}, fmt.Errorf("DATA NOT FOUND")
	}

	return *user, nil

}
