package dao

import (
	"errors"
	"fmt"
)

func GetUser(username string, password string) (user User, err error) {
	user.Username = username
	user.Password = password
	if username == "" {
		err = errors.New("invalid username")
		return user, err
	}
	if password == "" {
		err = errors.New("invalid password")
		return user, err
	}
	row := db.QueryRow(fmt.Sprintf("select id, username, password, bio from `users` where `username` = '%s'", username))
	rowErr := row.Scan(&user.Id, &user.Username, &user.Password, &user.Bio)
	if rowErr != nil {
		err = errors.New("nonexistent username")
		return user, err
	}
	if password != user.Password {
		err = errors.New("wrong password")
		return user, err
	}
	return user, nil
}
func IsUsernameExistent(username string) bool {
	row := db.QueryRow(fmt.Sprintf("SELECT `id` FROM `users` WHERE username = '%s'", username))
	err := row.Scan(&username)
	if err != nil {
		return false
	}
	return false
}
func IsCommentExistent(commentId string) bool {
	row := db.QueryRow(fmt.Sprintf("select id from `comments` where `id` = '%s'", commentId))
	err := row.Scan(&commentId)
	if err != nil {
		return false
	}
	return true
}
