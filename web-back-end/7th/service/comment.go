package service

import (
	"errors"
	"messageBoard/dao"
	"strconv"
	"strings"
)

func GetComments(userId string) ([]dao.Comment, error) {
	comments, err := dao.GetComments(userId)
	if err != nil {
		return comments, err
	}
	// 遍历评论添加子节点（回复）
	for i := range comments {
		// 判断评论是否存在子节点
		if comments[i].ChildrenId != "" {
			err := AppendChildrenComment(&comments[i], userId)
			if err != nil {
				return comments, err
			}
		}
	}
	return comments, nil
}
func AppendChildrenComment(comment *dao.Comment, userId string) error {
	childrenId := strings.Split(comment.ChildrenId, ",")
	for _, childId := range childrenId {
		child, err := dao.GetCommentById(childId)
		if err != nil {
			comment.Children = append(comment.Children, nil)
			return nil
		}
		// 判断子节点是否在用户查看范围内
		if child.SecretTarget == -1 || strconv.Itoa(child.SecretTarget) == userId {
			// 判断子节点是否仍存在子节点
			if child.ChildrenId != "" {
				err := AppendChildrenComment(&child, userId)
				if err != nil {
					return err
				}
			}
			comment.Children = append(comment.Children, &child)
		}
	}
	return nil
}
func AddComment(value string, userId int, secretTarget string) error {
	err := dao.AddComment(value, userId, secretTarget)
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
func DeleteComment(target string) error {
	comment, err := dao.GetCommentById(target)
	if err != nil {
		return nil // 查找不到 ID 说明该评论已被删除，所以返回 nil
	}
	if comment.ChildrenId != "" {
		childrenId := strings.Split(comment.ChildrenId, ",")
		for _, childId := range childrenId {
			err = DeleteComment(childId)
			if err != nil {
				return err
			}
			err = dao.DeleteComment(childId)
			if err != nil {
				return err
			}
		}
	}

	err = dao.DeleteComment(target)
	if err != nil {
		return err
	}
	return nil
}
