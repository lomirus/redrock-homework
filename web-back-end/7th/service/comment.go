package service

import (
	"errors"
	"fmt"
	"messageBoard/dao"
)

func ShowComments() []dao.Comment {
	comments := dao.GetComments()
	for i := range comments {
		if comments[i].ParentId != -1 {
			comments[i].Value = fmt.Sprintf("%s -> %s", comments[i].Value, dao.GetParentComment(comments[i]))
		}
	}
	return comments
}
func AddComment(value string, userId int) error {
	err := dao.AddComment(value, userId)
	if err != nil {
		return err
	}
	return nil
}
func ReplyComment(target string, value string, userId int) error {
	if !dao.IsCommentExistent(target) {
		return errors.New("nonexistent target id")
	}
	err := dao.ReplyComment(target, value, userId)
	if err != nil {
		return err
	}
	return nil
}
func LikeComment(target string) error {
	err := dao.LikeComment(target)
	if err != nil {
		return err
	}
	return nil
}
