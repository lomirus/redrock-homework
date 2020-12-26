package model

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	Id       int
	Username string
	Password string
	Money    int
}

type Log struct {
	Id          int
	Username    string
	MoneySource string
	Remark      string
	Money       int
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:0@/test?charset=utf8")
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `users`(" +
		"id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY," +
		"username VARCHAR(32) NOT NULL," +
		"password VARCHAR(32) NOT NULL," +
		"money BIGINT NOT NULL DEFAULT 0" +
		") ENGINE = InnoDB CHARACTER SET utf8 COLLATE utf8_general_ci;")
	if err != nil {
		log.Fatal("Creating users Table: ", err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `logs`(" +
		"id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY," +
		"username VARCHAR(32) NOT NULL," +
		"money_source VARCHAR(32) NOT NULL," +
		"remark VARCHAR(256) NOT NULL DEFAULT ''," +
		"money BIGINT NOT NULL" +
		") ENGINE = InnoDB CHARACTER SET utf8 COLLATE utf8_general_ci;")
	if err != nil {
		log.Fatal("Creating users Table: ", err)
	}
}

func Register(username string, password string) (err error) {
	var user User
	row := db.QueryRow("select `id` from users where username=?", username)
	err = row.Scan(&user.Id)
	if err == nil {
		return errors.New("the username has been registered")
	}
	_, err = db.Exec("INSERT INTO `users` (`username`,`password`) VALUES (?,?)", username, password)
	return err
}
func Login(username string, password string) (err error) {
	var user User
	row := db.QueryRow("select `id` from users where username=?", username)
	err = row.Scan(&user.Id)
	if err != nil {
		return errors.New("cannot find the username")
	}
	row = db.QueryRow("select `id` from users where username=? and password=?", username, password)
	err = row.Scan(&user.Id)
	if err != nil {
		return errors.New("wrong password")
	}
	return nil
}
func GetUser(username string) (user User, err error) {
	row := db.QueryRow("select `id`,`username`,`password`,`money` from users where username=?", username)
	err = row.Scan(&user.Id, &user.Username, &user.Password, &user.Money)
	if err != nil {
		return user, errors.New("cannot find the user")
	}
	return user, nil
}
func Charge(self User, money int) (err error) {
	self.Money += money
	_, err = db.Exec("update `users` set money=? where username=?", self.Money, self.Username)
	if err != nil {
		return err
	}
	_, err = db.Exec("insert into `logs` (username, money_source, money) values (?,?,?)",
		self.Username, self.Username, money)
	if err != nil {
		return err
	}
	return nil
}
func Transfer(self User, target User, money int, remark string) (err error) {
	self.Money -= money
	target.Money += money
	_, err = db.Exec("update `users` set money=? where username=?", self.Money, self.Username)
	if err != nil {
		return err
	}
	_, err = db.Exec("update `users` set money=? where username=?", target.Money, target.Username)
	if err != nil {
		return err
	}
	_, err = db.Exec("insert into `logs` (username, money_source, money, remark) values (?,?,?,?)",
		self.Username, self.Username, -money, remark)
	if err != nil {
		return err
	}
	_, err = db.Exec("insert into `logs` (username, money_source, money, remark) values (?,?,?,?)",
		target.Username, self.Username, money, remark)
	if err != nil {
		return err
	}
	return nil
}
func GetLogs(username string) (logs []Log, err error) {
	rows, err := db.Query("select `id`,`username`,`money_source`,`remark`,`money` from `logs` where `username`=?", username)
	if err != nil {
		return logs, err
	}
	for rows.Next() {
		var log Log
		err = rows.Scan(&log.Id, &log.Username, &log.MoneySource, &log.Remark, &log.Money)
		if err != nil {
			return logs, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}
