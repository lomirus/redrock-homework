package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)
import _ "github.com/go-sql-driver/mysql"

type User struct {
	Id int
	Username string
	Password string
	Bio string
}
type Comment struct {
	Id       int
	ParentId int
	Value string
}

var db *sql.DB

func init(){
	var err error
	db, err = sql.Open("mysql", "root:0@/test?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("CREATE TABLE IF NOT EXISTS `users`(" +
		"id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY," +
		"username VARCHAR(32) NOT NULL," +
		"password VARCHAR(32) NOT NULL," +
		"bio VARCHAR(128) DEFAULT ''" +
		") ENGINE = InnoDB CHARACTER SET utf8 COLLATE utf8_general_ci;")
	db.Exec("CREATE TABLE IF NOT EXISTS `comments`(" +
		"id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY," +
		"parent_id BIGINT NOT NULL DEFAULT -1," +
		"value VARCHAR(256) NOT NULL DEFAULT ''" +
		") ENGINE = InnoDB CHARACTER SET utf8 COLLATE utf8_general_ci;")
}

func main() {
	defer db.Close()

	for {
		var input string
		fmt.Println("Input 1 to login;")
		fmt.Println("Input 2 to register;")
		fmt.Println("Input 3 to exit;")
		fmt.Scan(&input)
		switch input {
		case "1":{
			fmt.Println("LOGIN")
			login(db)
		}
		case "2":{
			fmt.Println("REGISTER")
			register(db)
		}
		case "3":{
			os.Exit(1)
		}
		default:{
			fmt.Println("Error: Unexpected Input")
		}
		}
	}
}

func login(db *sql.DB){
	for {
		var username string
		var password string
		fmt.Println("Please input your username: ")
		fmt.Scan(&username)
		fmt.Println("Please input your password: ")
		fmt.Scan(&password)
		var user User
		row := db.QueryRow(fmt.Sprintf("select id, username, password, bio from `users` where `username` = '%s'", username))
		err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Bio)
		if err != nil {
			fmt.Println("Error: Nonexistent Password")
			return
		}
		if password != user.Password{
			fmt.Println("Error: Wrong Password")
			return
		}
		fmt.Println("Info: Login Successful")
		onLogin(user)
		return
	}
}

func register(db *sql.DB){
	for {
		var username string
		var password string
		var confirm string
		fmt.Println("Please input your username: ")
		fmt.Scan(&username)
		fmt.Println("Please input your password: ")
		fmt.Scan(&password)
		fmt.Println("Please confirm your password: ")
		fmt.Scan(&confirm)
		if password != confirm {
			fmt.Println("Error: Inconsistent Password")
			return
		}
		query, err := db.Query(fmt.Sprintf("SELECT `username` FROM `users` where `username` = '%s'", username))
		if err != nil {
			log.Fatal(err)
		}
		if query.Next() {
			fmt.Println("Error: Duplicate Username")
			return
		}
		_, err = db.Exec("INSERT INTO `users` (`username`,`password`) VALUES (?,?)", username, password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Info: Registration Successful")
		return
	}
}

func onLogin(user User){
	fmt.Printf("Your UID is %d;\n", user.Id)
	fmt.Printf("Your username is %s;\n", user.Username)
	if user.Bio == "" {
		fmt.Println("Your do not have a bio yet;")
	} else {
		fmt.Printf("Your bio is %s;\n", user.Bio)
	}
	for {
		var input string
		fmt.Println("Input 1 to edit your username;")
		fmt.Println("Input 2 to edit your password;")
		fmt.Println("Input 3 to edit your bio;")
		fmt.Println("Input 4 to show comments;")
		fmt.Println("Input 5 to add a comment;")
		fmt.Println("Input 6 to reply a comment;")
		fmt.Println("Input 7 to log out;")
		fmt.Scan(&input)
		switch input {
		case "1":{
			var newUsername string
			fmt.Println("Please input your new username;")
			fmt.Scan(&newUsername)
			if newUsername == user.Username {
				fmt.Println("Warning: Invalid Operation")
				continue
			}
			query, err := db.Query(fmt.Sprintf("SELECT `id` FROM `users` WHERE username = '%s'", newUsername))
			if err != nil {
				log.Fatal(err)
			}
			if query.Next() {
				fmt.Println("Error: Duplicate Username")
				continue
			}
			user.Username = newUsername
			_, err = db.Exec("UPDATE `users` SET `username` = ? WHERE `id` = ?", user.Username, user.Id)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Info: Updated Username Successfully")
		}
		case "2":{
			var newPassword string
			var cfmPassword string
			fmt.Println("Please input your new password;")
			fmt.Scan(&newPassword)
			if newPassword == user.Password {
				fmt.Println("Warning: Invalid Operation")
				continue
			}
			fmt.Println("Please confirm your new password;")
			fmt.Scan(&cfmPassword)
			if newPassword == cfmPassword {
				fmt.Println("Error: Inconsistent Password")
				continue
			}

			user.Password = newPassword
			_, err := db.Exec("UPDATE `users` SET `password` = ? WHERE `id` = ?", user.Password, user.Id)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Info: Updated Password Successfully")
		}
		case "3":{
			var newBio string
			fmt.Println("Please input your new bio;")
			fmt.Scan(&newBio)
			if newBio == user.Bio {
				fmt.Println("Warning: Invalid Operation")
				continue
			}
			user.Bio = newBio
			_, err := db.Exec("UPDATE `users` SET `bio` = ? WHERE `id` = ?", user.Bio, user.Id)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Info: Updated Bio Successfully")
		}
		case "4":{// to show comments
			comments := getComments(db)
			for _, v := range comments {
				if v.ParentId != -1 {
					v.Value = fmt.Sprintf("%s -> %s", v.Value, getParentComment(v))
				}
				fmt.Printf("ID:%d\tPrtID:%d \t%s\n", v.Id, v.ParentId, v.Value)
			}
		}
		case "5":{// to add a comment
			var input string
			fmt.Println("Please input your comment:")
			fmt.Scanln(&input)
			_, err := db.Exec("INSERT INTO `comments` (`value`) VALUES (?)", input)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Info: Commented Successfully")
		}
		case "6":{// to reply a comment
			var replyToId string
			var reply string
			fmt.Println("Please input the id you are going to reply:")
			fmt.Scanln(&replyToId)
			fmt.Println("Please input your reply:")
			fmt.Scanln(&reply)
			isIdExistent := db.QueryRow(fmt.Sprintf("select id from `comments` where `id` = '%s'", replyToId))
			err := isIdExistent.Scan(&replyToId)
			if err != nil {
				fmt.Println("Error: Nonexistent Comment Id")
				continue
			}
			_, err = db.Exec("INSERT INTO `comments` (`value`, `parent_id`) VALUES (?,?)", reply, replyToId)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Info: Replied Successfully")
		}
		case "7":{
			fmt.Println("Info: Logged Out Successfully")
			return
		}
		default:
			fmt.Println("Error: Unexpected Input")
		}
	}
}

func getComments(db *sql.DB) []Comment {
	var comments []Comment
	query, err := db.Query("SELECT `id`, `parent_id`, `value` FROM `comments`")
	if err != nil {
		log.Fatal(err)
	}
	for query.Next() {
		var newComment Comment
		query.Scan(&newComment.Id, &newComment.ParentId, &newComment.Value)
		comments = append(comments, newComment)
	}
	return comments
}
func addComment(){

}
func rplComment(){

}
func getParentComment(comment Comment) string {
	var parentComment Comment
	row := db.QueryRow(fmt.Sprintf("select `id`,`parent_id`,`value` from `comments` where `id` = '%d'", comment.ParentId))
	err := row.Scan(&parentComment.Id, &parentComment.ParentId, &parentComment.Value)
	if err != nil {
		log.Fatal(err)
	}
	if parentComment.ParentId != -1 {
		parentComment.Value = fmt.Sprintf("%s -> %s", parentComment.Value, getParentComment(parentComment))
	}
	return parentComment.Value
}
