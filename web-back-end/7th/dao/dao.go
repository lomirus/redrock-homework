package dao

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}
type Comment struct {
	Id           int    `json:"id"`
	ParentId     int    `json:"parent_id"`
	ChildrenId   string `json:"children_id"`
	UserId       int    `json:"user_id"`
	SecretTarget int    `json:"secret_target"`
	Value        string `json:"value"`
	Likes        int    `json:"likes"`
	Children     []*Comment
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:0@/test?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `users`(" +
		"id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY," +
		"username VARCHAR(32) NOT NULL," +
		"password VARCHAR(32) NOT NULL," +
		"bio VARCHAR(128) DEFAULT ''" +
		") ENGINE = InnoDB CHARACTER SET utf8 COLLATE utf8_general_ci;")
	if err != nil {
		log.Fatal("Creating users Table: ", err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `comments`(" +
		"id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY," +
		"user_id BIGINT NOT NULL," +
		"parent_id BIGINT NOT NULL DEFAULT -1," +
		"children_id VARCHAR(1024) NOT NULL DEFAULT ''," +
		"secret_target BIGINT NOT NULL DEFAULT -1," +
		"value VARCHAR(256) NOT NULL DEFAULT ''," +
		"likes BIGINT NOT NULL DEFAULT 0" +
		") ENGINE = InnoDB CHARACTER SET utf8 COLLATE utf8_general_ci;")
	if err != nil {
		log.Fatal("Creating comments Table: ", err)
	}
}

func Register(username string, password string) error {
	_, err := db.Exec("INSERT INTO `users` (`username`,`password`) VALUES (?,?)", username, password)
	if err == nil {
		return err
	}
	return nil
}

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
func GetUserById(userId int) (user User, err error) {
	row := db.QueryRow(fmt.Sprintf("select id, username, password, bio from `users` where `id` = '%d'", userId))
	rowErr := row.Scan(&user.Id, &user.Username, &user.Password, &user.Bio)
	if rowErr != nil {
		err = errors.New("the id of user does not exist")
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
