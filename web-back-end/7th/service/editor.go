package service

import (
	"errors"
	"messageBoard/dao"
)

func EditUsername(userId int, newUsername string) error {
	user, err := dao.GetUserById(userId)
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
func EditPassword(userId int, newPassword string) error {
	user, err := dao.GetUserById(userId)
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
func EditBio(userId int, newBio string) error {
	user, err := dao.GetUserById(userId)
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
