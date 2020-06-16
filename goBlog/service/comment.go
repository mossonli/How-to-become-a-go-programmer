package service

import (
	"fmt"
	"goBlog/dao"
	"goBlog/model"
	"time"
)

func GetCommentList(articleId int64) (commentList []*model.Comment, err error) {
	exist, err := dao.IsArticleExist(articleId)
	if err != nil {
		fmt.Printf("query database failed, err:%v\n", err)
		return
	}
	if exist == false {
		err = fmt.Errorf("article id:%d not found", articleId)
		return
	}
	commentList, err = dao.GetCommentList(articleId, 0, 100)
	return
}

// 插入评论
func InsertComment(comment,author,email string, articleId int64) (err error) {
	exist, err := dao.IsArticleExist(articleId)
	if err != nil {
		fmt.Printf("query database failed, err:%v\n", err)
		return
	}
	if exist == false {
		err = fmt.Errorf("article id:%d not found", articleId)
		return
	}
	var c model.Comment
	c.ArticleId = articleId
	c.Content = comment
	c.Username = author
	c.CreateTime = time.Now()
	c.Status = 1
	err = dao.InsertComment(&c)
	return
}