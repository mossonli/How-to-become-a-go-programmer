package dao

import (
	"fmt"
	"goBlog/model"
)
// 插入评论
func InsertComment(comment *model.Comment) (err error) {
	if comment == nil {
		err = fmt.Errorf("invalid parameter")
		return
	}
	// 添加事物
	tx, err := DB.Beginx()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()
	sqlstr := `insert into 
					comment(content, username, article_id)
				values (?, ?, ?)`
	_, err = tx.Exec(sqlstr, comment.Content, comment.Username, comment.ArticleId)
	if err != nil {
		return
	}

	sqlstr = `  update article 
				set 
					comment_count = comment_count + 1
				where
					id = ?`
	_, err = tx.Exec(sqlstr, comment.ArticleId)
	if err != nil {
		return
	}
	err = tx.Commit() // 提交事物
	return
}

// 更新评论数
func UpdateViewCount(articleId int64) (err error) {
	sqlstr := ` update article 
				set 
					comment_count = comment_count + 1
				where
					id = ?`

	_, err = DB.Exec(sqlstr, articleId)
	if err != nil {
		return
	}
	return
}

// 获取评论列表
func GetCommentList(articleId int64, pageNum, pageSize int) (commentList []*model.Comment, err error) {
	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d, page_size:%d", pageNum, pageSize)
		return
	}
	sqlStr := `select 
					id, content, username, create_time
				from 
					comment 
				where 
					article_id = ? and 
					status = 1
				order by create_time desc
				limit ?, ?`

	err = DB.Select(&commentList, sqlStr, articleId, pageNum, pageSize)
	return
}