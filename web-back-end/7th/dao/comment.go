package dao

import (
	"fmt"
	"strconv"
)

func GetComments() ([]Comment, error) {
	var comments []Comment
	query, err := db.Query("SELECT `id`, `user_id`,`parent_id`,`children_id`, `value`, `likes`" +
		" FROM `comments` WHERE parent_id=-1")
	if err != nil {
		return comments, err
	}
	for query.Next() {
		var newComment Comment
		err := query.Scan(&newComment.Id, &newComment.UserId, &newComment.ParentId, &newComment.ChildrenId,
			&newComment.Value, &newComment.Likes)
		if err != nil {
			return comments, err
		}
		comments = append(comments, newComment)
	}
	return comments, nil
}
func GetCommentById(id string) (Comment, error) {
	var comment Comment
	row := db.QueryRow("SELECT `id`, `user_id`,`parent_id`,`children_id`, `value`, `likes`"+
		" FROM `comments` WHERE id=?", id)
	err := row.Scan(&comment.Id, &comment.UserId, &comment.ParentId, &comment.ChildrenId,
		&comment.Value, &comment.Likes)
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func AddComment(value string, userId int) error {
	_, err := db.Exec("INSERT INTO `comments` (`value`,`user_id`) VALUES (?,?)", value, userId)
	if err != nil {
		return err
	}
	return nil
}
func ReplyComment(target string, value string, userId int) error {
	// 插入回复
	reply, err := db.Exec("INSERT INTO `comments` (`value`, `parent_id`, `user_id`) VALUES (?,?,?)",
		value, target, userId)
	if err != nil {
		return err
	}
	lastInsertId, err := reply.LastInsertId()
	if err != nil {
		return err
	}
	// 获取父元素的子节点ID
	var childrenId string
	row := db.QueryRow("SELECT `children_id` FROM `comments` WHERE `id`=?", target)
	err = row.Scan(&childrenId)
	if err != nil {
		return err
	}
	// 更新子节点ID
	if childrenId == "" {
		childrenId += strconv.FormatInt(lastInsertId, 10)
	} else {
		childrenId += fmt.Sprintf("%s%d", ",", lastInsertId)
	}
	// 更新父元素的子节点ID
	_, err = db.Exec("UPDATE `comments` SET `children_id`=? WHERE `id`=?", childrenId, target)
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
