package service

import (
	"entity"
	"fmt"
)

// Login logs in with provided user
func Login(username string, password string) error {
	if len(entity.CurSessionModel.GetCurUser()) > 0 {
		return fmt.Errorf("you have logged in as '%s'", entity.CurSessionModel.GetCurUser())
	}
	user := entity.UserModel.FindByUsername(username)
	if user != nil && MatchPassword(user, password) == true {
		entity.CurSessionModel.SetCurUser(user)
	} else {
		return fmt.Errorf("username or password is incorrent")
	}
	return nil
}

func hasLoggedin() error {
	if len(entity.CurSessionModel.GetCurUser()) == 0 {
		return fmt.Errorf("you haven't login")
	}
	return nil
}

// Logout logs out
func Logout() error {
	var err error
	err = hasLoggedin()
	if err == nil {
		entity.CurSessionModel.SetCurUser(&user{})
	} else {
		return err
	}
	return nil
}
