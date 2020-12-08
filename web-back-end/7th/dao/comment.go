package dao

import (
	"fmt"
	"log"
)

func GetComments() []Comment {
	var comments []Comment
	query, err := db.Query("SELECT `id`, `user_id`,`parent_id`, `value`, `likes` FROM `comments`")
	if err != nil {
		log.Fatal(err)
	}
	for query.Next() {
		var newComment Comment
		query.Scan(&newComment.Id, &newComment.UserId, &newComment.ParentId,
			&newComment.Value, &newComment.Likes)
		comments = append(comments, newComment)
	}
	return comments
}
func GetParentComment(comment Comment) string {
	var parentComment Comment
	row := db.QueryRow(fmt.Sprintf("select `id`,`user_id`,`parent_id`,`value` from `comments` where `id` = '%d'", comment.ParentId))
	err := row.Scan(&parentComment.Id, &parentComment.UserId, &parentComment.ParentId, &parentComment.Value)
	if err != nil {
		log.Fatal(err)
	}
	if parentComment.ParentId != -1 {
		parentComment.Value = fmt.Sprintf("%s -> %s", parentComment.Value, GetParentComment(parentComment))
	}
	return parentComment.Value
}

func AddComment(value string, userId int) error {
	_, err := db.Exec("INSERT INTO `comments` (`value`,`user_id`) VALUES (?,?)", value, userId)
	if err != nil {
		return err
	}
	return nil
}
func ReplyComment(target string, value string, userId int) error {
	_, err := db.Exec("INSERT INTO `comments` (`value`, `parent_id`, `user_id`) VALUES (?,?,?)",
		value, target, userId)
	if err != nil {
		return err
	}
	return nil
}
func LikeComment(target string) error {
	var oldLikes int
	row := db.QueryRow("SELECT `likes` FROM `comments` WHERE `id`=?", target)
	err := row.Scan(&oldLikes)
	if err != nil {
		return err
	}
	newLikes := oldLikes + 1
	_, err = db.Exec("UPDATE `comments` SET `likes`=? WHERE `id`=?", newLikes, target)
	if err != nil {
		return err
	}
	return nil
}
