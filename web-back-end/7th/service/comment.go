package service

import (
	"errors"
	"messageBoard/dao"
	"strings"
)

func GetComments() ([]dao.Comment, error) {
	comments, err := dao.GetComments()
	if err != nil {
		return comments, err
	}
	// 遍历评论添加子节点（回复）
	for i := range comments {
		// 判断评论是否存在子节点
		if comments[i].ChildrenId != "" {
			err := AppendChildrenComment(&comments[i])
			if err != nil {
				return comments, err
			}
		}
	}
	return comments, nil
}
func AppendChildrenComment(comment *dao.Comment) error {
	childrenId := strings.Split(comment.ChildrenId, ",")
	for _, childId := range childrenId {
		child, err := dao.GetCommentById(childId)
		if err != nil {
			return err
		}
		// 判断子节点是否仍存在子节点
		if child.ChildrenId != "" {
			err := AppendChildrenComment(&child)
			if err != nil {
				return err
			}
		}
		comment.Children = append(comment.Children, &child)
	}
	return nil
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
