package service

import (
	"errors"
	"messageBoard/dao"
)

func EditUsername(username string, password string, newUsername string) error {
	user, err := dao.GetUser(username, password)
	if err != nil {
		return err
	}
	if newUsername == user.Username {
		return errors.New("invalid operation")
	}
	if dao.IsUsernameExistent(newUsername) {
		return errors.New("existed username")
	}
	err = dao.EditUsername(user.Id, newUsername)
	if err != nil {
		return err
	}
	return nil
}
func EditPassword(username string, password string, newPassword string) error {
	user, err := dao.GetUser(username, password)
	if err != nil {
		return err
	}
	if newPassword == user.Password {
		return errors.New("invalid operation")
	}
	err = dao.EditPassword(user.Id, newPassword)
	if err != nil {
		return err
	}
	return nil
}
func EditBio(username string, password string, newBio string) error {
	user, err := dao.GetUser(username, password)
	if err != nil {
		return err
	}
	if newBio == user.Bio {
		return errors.New("invalid operation")
	}
	err = dao.EditBio(user.Id, newBio)
	if err != nil {
		return err
	}
	return nil
}
