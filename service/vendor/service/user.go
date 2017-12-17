package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"utils"

	"entities"

	"github.com/unrolled/render"
)

func deleteAccHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := checkIsLogin(r)
		// TODO: delete meetings
		// ...
		entities.UserServ.Delete(user)
		logoutHandler(formatter)(w, r)
	}
}

func registerHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user entities.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		var badData struct {
			Msg string `json:"msg"`
		}

		if err := newUservalidate(&user); err != nil {
			badData.Msg = err.Error()
			formatter.JSON(w, http.StatusBadRequest, badData)
			return
		}

		user.Password = utils.Sha(user.Password)
		entities.UserServ.Add(&user)
		formatter.JSON(w, http.StatusCreated, struct{}{})
	}
}

func newUservalidate(user *entities.User) error {
	if len(user.Username) == 0 {
		return fmt.Errorf("username should not be empty")
	}
	if entities.UserServ.FindByUsername(user.Username) != nil {
		return fmt.Errorf("username '%s' already exists", user.Username)
	}
	if len(user.Password) == 0 {
		return fmt.Errorf("password should not be empty")
	}
	// ...TODO
	return nil
}

func findAllUsersHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		formatter.JSON(w, http.StatusOK, entities.UserServ.FindAll())
	}
}
