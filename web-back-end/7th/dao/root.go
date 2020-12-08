package dao

import (
	"database/sql"
	"log"
)
import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}
type Comment struct {
	Id       int    `json:"id"`
	UserId   int    `json:"user_id"`
	ParentId int    `json:"parent_id"`
	Value    string `json:"value"`
	Likes    string `json:"likes"`
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
		log.Fatal(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `comments`(" +
		"id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY," +
		"user_id BIGINT NOT NULL," +
		"parent_id BIGINT NOT NULL DEFAULT -1," +
		"value VARCHAR(256) NOT NULL DEFAULT ''" +
		") ENGINE = InnoDB CHARACTER SET utf8 COLLATE utf8_general_ci;")
	if err != nil {
		log.Fatal(err)
	}
}

func Register(username string, password string) error {
	_, err := db.Exec("INSERT INTO `users` (`username`,`password`) VALUES (?,?)", username, password)
	if err == nil {
		return err
	}
	return nil
}
