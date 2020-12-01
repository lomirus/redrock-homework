package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)
import _ "github.com/go-sql-driver/mysql"
import "github.com/gin-gonic/gin"

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
}

var db *sql.DB

func init() {
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
		"user_id BIGINT NOT NULL," +
		"parent_id BIGINT NOT NULL DEFAULT -1," +
		"value VARCHAR(256) NOT NULL DEFAULT ''" +
		") ENGINE = InnoDB CHARACTER SET utf8 COLLATE utf8_general_ci;")
}

func main() {
	defer db.Close()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		var user User
		var err error
		user.Username, err = c.Cookie("username")
		if err != nil {
			c.String(http.StatusOK, "Please Login at '/login'")
			return
		}
		user.Password, _ = c.Cookie("password")
		strId, _ := c.Cookie("id")
		user.Id, _ = strconv.Atoi(strId)
		user.Bio, _ = c.Cookie("bio")
		c.String(http.StatusOK, fmt.Sprintf("Hello, %s!\nYour ID is %d;\nYour Bio is %s.",
			user.Username, user.Id, user.Bio))
	})
	router.GET("/login", func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")
		if username == "" {
			c.JSON(http.StatusOK, gin.H{
				"Error": "Invalid Username",
			})
			return
		}
		if password == "" {
			c.JSON(http.StatusOK, gin.H{
				"Error": "Invalid Password",
			})
			return
		}
		var user User
		row := db.QueryRow(fmt.Sprintf("select id, username, password, bio from `users` where `username` = '%s'", username))
		err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Bio)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Error": "Nonexistent Username",
			})
			return
		}
		if password != user.Password {
			c.JSON(http.StatusOK, gin.H{
				"Error": "Wrong Password",
			})
			return
		}
		c.SetCookie("username", user.Username, 3600, "/", "localhost", false, true)
		c.SetCookie("password", user.Password, 3600, "/", "localhost", false, true)
		c.SetCookie("id", fmt.Sprintf("%d", user.Id), 3600, "/", "localhost", false, true)
		c.SetCookie("bio", user.Bio, 3600, "/", "localhost", false, true)
		c.String(http.StatusOK, "Login Successfully")
	})
	router.GET("/register", func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")
		query, err := db.Query(fmt.Sprintf("SELECT `username` FROM `users` where `username` = '%s'", username))
		if err != nil {
			log.Println(err)
			return
		}
		if query.Next() {
			c.JSON(http.StatusOK, gin.H{
				"Error": "Duplicate Username",
			})
			return
		}
		_, err = db.Exec("INSERT INTO `users` (`username`,`password`) VALUES (?,?)", username, password)
		if err != nil {
			log.Println(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Info": "Registration Successful",
		})
		return
		fmt.Println("REGISTER")
	})
	router.GET("/logout", func(c *gin.Context) {
		c.SetCookie("username", "", -1, "/", "localhost", false, true)
		c.SetCookie("password", "", -1, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{
			"Info": "Logged out Successfully",
		})
	})
	editGroup := router.Group("edit", verify)
	{
		editGroup.GET("/username", func(c *gin.Context) {
			username, _ := c.Cookie("username")
			password, _ := c.Cookie("username")
			user, _ := getUser(username, password)
			newUsername := c.Query("new")
			if newUsername == user.Username {
				c.JSON(http.StatusOK, gin.H{
					"Warning": "Invalid Operation",
				})
				return
			}
			query, err := db.Query(fmt.Sprintf("SELECT `id` FROM `users` WHERE username = '%s'", newUsername))
			if err != nil {
				log.Println(err)
				return
			}
			if query.Next() {
				c.JSON(http.StatusOK, gin.H{
					"Warning": "Duplicate Username",
				})
				return
			}
			user.Username = newUsername
			_, err = db.Exec("UPDATE `users` SET `username` = ? WHERE `id` = ?", user.Username, user.Id)
			if err != nil {
				log.Println(err)
				return
			}
			c.SetCookie("username", user.Username, 3600, "/", "localhost", false, true)
			c.JSON(http.StatusOK, gin.H{
				"Info": "Updated Username Successfully",
			})
		})
		editGroup.GET("/password", func(c *gin.Context) {
			username, _ := c.Cookie("username")
			password, _ := c.Cookie("username")
			user, _ := getUser(username, password)
			newPassword := c.Query("new")
			if newPassword == user.Password {
				c.JSON(http.StatusOK, gin.H{
					"Warning": "Invalid Operation",
				})
				return
			}
			user.Password = newPassword
			_, err := db.Exec("UPDATE `users` SET `password` = ? WHERE `id` = ?", user.Password, user.Id)
			if err != nil {
				log.Println(err)
				return
			}
			c.SetCookie("password", user.Password, 3600, "/", "localhost", false, true)
			c.JSON(http.StatusOK, gin.H{
				"Info": "Updated Password Successfully",
			})
		})
		editGroup.GET("/bio", func(c *gin.Context) {
			username, _ := c.Cookie("username")
			password, _ := c.Cookie("username")
			user, _ := getUser(username, password)
			newBio := c.Query("new")
			if newBio == user.Bio {
				c.JSON(http.StatusOK, gin.H{
					"Warning": "Invalid Operation",
				})
				return
			}
			user.Bio = newBio
			_, err := db.Exec("UPDATE `users` SET `bio` = ? WHERE `id` = ?", user.Bio, user.Id)
			if err != nil {
				log.Println(err)
				return
			}
			c.SetCookie("bio", user.Bio, 3600, "/", "localhost", false, true)
			c.JSON(http.StatusOK, gin.H{
				"Info": "Updated Bio Successfully",
			})
		})
	}
	commentGroup := router.Group("comment", verify)
	{
		commentGroup.GET("show", func(c *gin.Context) {
			comments := getComments(db)
			for i := range comments {
				if comments[i].ParentId != -1 {
					comments[i].Value = fmt.Sprintf("%s -> %s", comments[i].Value, getParentComment(comments[i]))
				}
			}
			c.JSON(http.StatusOK, comments)
		})
		commentGroup.GET("add", func(c *gin.Context) {
			value := c.Query("value")
			userId, _ := c.Cookie("id")
			if value == "" {
				c.JSON(http.StatusOK, gin.H{
					"Error": "Invalid Comment Value",
				})
				return
			}
			_, err := db.Exec("INSERT INTO `comments` (`value`,`user_id`) VALUES (?,?)", value, userId)
			if err != nil {
				log.Println(err)
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"Info": "Commented Successfully",
			})
		})
		commentGroup.GET("reply", func(c *gin.Context) {
			target := c.Query("target")
			value := c.Query("value")
			userId, _ := c.Cookie("id")
			if value == "" {
				c.JSON(http.StatusOK, gin.H{
					"Error": "Invalid Reply Value",
				})
				return
			}
			isIdExistent := db.QueryRow(fmt.Sprintf("select id from `comments` where `id` = '%s'", target))
			err := isIdExistent.Scan(&target)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"Error": "Nonexistent Target ID",
				})
				return
			}
			_, err = db.Exec("INSERT INTO `comments` (`value`, `parent_id`, `user_id`) VALUES (?,?,?)",
				value, target, userId)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"Error": err,
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"Info": "Replied Successfully",
			})
		})
	}
	router.Run()
}

func getUser(username string, password string) (user User, err error) {
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
	return user, err
}
func verify(c *gin.Context) {
	username, err := c.Cookie("username")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Error": "Cannot find username, please login first",
		})
		c.Abort()
		return
	}
	password, err := c.Cookie("password")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Error": "Cannot find password, please login first",
		})
		c.Abort()
		return
	}
	_, err = getUser(username, password)
	if err == nil {
		c.Next()
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Error": "Username or password is wrong, please login again",
		})
		c.Abort()
		return
	}
}
func getComments(db *sql.DB) []Comment {
	var comments []Comment
	query, err := db.Query("SELECT `id`, `user_id`,`parent_id`, `value` FROM `comments`")
	if err != nil {
		log.Fatal(err)
	}
	for query.Next() {
		var newComment Comment
		query.Scan(&newComment.Id, &newComment.UserId, &newComment.ParentId, &newComment.Value)
		comments = append(comments, newComment)
	}
	return comments
}
func getParentComment(comment Comment) string {
	var parentComment Comment
	row := db.QueryRow(fmt.Sprintf("select `id`,`user_id`,`parent_id`,`value` from `comments` where `id` = '%d'", comment.ParentId))
	err := row.Scan(&parentComment.Id, &parentComment.UserId, &parentComment.ParentId, &parentComment.Value)
	if err != nil {
		log.Fatal(err)
	}
	if parentComment.ParentId != -1 {
		parentComment.Value = fmt.Sprintf("%s -> %s", parentComment.Value, getParentComment(parentComment))
	}
	return parentComment.Value
}
