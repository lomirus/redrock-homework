package service

import (
	"errors"
	"messageBoard/dao"
)

func Login(username string, password string) (dao.User, error) {
	user, err := dao.GetUser(username, password)
	if err != nil {
		return user, err
	}
	return user, nil
}
func Register(username string, password string) error {
	if dao.IsUsernameExistent(username) {
		return errors.New("existed username")
	}
	err := dao.Register(username, password)
	if err != nil {
		return err
	}
	return nil
}
