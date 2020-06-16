package service

import (
	"fmt"
	"goBlog/dao"
	"goBlog/model"
)

// 获取评论列表
func GetLeaveList() (leaveList []*model.Leave, err error) {
	leaveList, err = dao.GetLeaveList()
	if err != nil {
		fmt.Printf("get leave list failed, err:%v\n", err)
		return
	}
	return
}

// 插入新评论
func InsertLeave(username, email, content string) (err error) {
	var leave model.Leave
	leave.Content = content
	leave.Email = email
	leave.Username = username
	err = dao.InsertLeave(&leave)
	if err != nil {
		fmt.Printf("insert leave failed, err:%v, leave:%#v\n", err, leave)
		return
	}
	return
}