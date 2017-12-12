package service

import (
	"entity"
	"fmt"
	"utils"
)

type user = entity.User

var (
	// ErrInvalidUsername is error for incorrent username
	ErrInvalidUsername = fmt.Errorf("username should not be empty")
	// ErrInvalidPassword is error for incorrentpassword
	ErrInvalidPassword = fmt.Errorf("password should not be empty")
	// ErrRepeatUsername is error for incorrentpassword
	ErrRepeatUsername = fmt.Errorf("you input username already exists")
)

func validateNewUser(user *user) error {
	if len(user.Username) == 0 {
		return ErrInvalidUsername
	}
	if len(user.Password) == 0 {
		return ErrInvalidPassword
	}
	if entity.UserModel.FindByUsername(user.Username) != nil {
		return ErrRepeatUsername
	}
	return nil
}

// Register an agenda account
func Register(username string, password string, email string, telephone string) (err error) {
	newUser := &user{
		Username:  username,
		Password:  password,
		Email:     email,
		Telephone: telephone,
	}
	err = validateNewUser(newUser)
	if err != nil {
		return
	}
	// use the jiami algorithm sha1 to store the password
	newUser.Password = utils.Sha(newUser.Password)
	entity.UserModel.AddUser(newUser)
	return
}

// GetAllUsers find all registered users
func GetAllUsers() ([]user, error) {
	if len(entity.CurSessionModel.GetCurUser()) == 0 {
		return nil, fmt.Errorf("You haven't logined")
	}
	users := entity.UserModel.FindBy(func(u *user) bool {
		return true
	})
	return users, nil
}

// MatchPassword : judge whether the password is correct
func MatchPassword(user *user, password string) bool {
	return utils.Sha(password) == user.Password
}
