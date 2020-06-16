package dao

import (
	"database/sql"
	"fmt"
	"goBlog/model"
)

// 文章插入
func InsertArticle(article *model.ArticleDetail) (articleId int64, err error) {
	result, err := DB.Exec("insert into article(content,summary,title,username,category_id, view_count,comment_count)"+
		"values(?,?,?,?,?,?,?)", article.Content, article.Summary, article.Title, article.Username,article.ArticleInfo.CategoryId,
		article.ArticleInfo.ViewCount, article.ArticleInfo.CommentCount)
	if err != nil {
		fmt.Println(err)
		return
	}
	articleId, err = result.LastInsertId()
	return
}

// 分页获取所有的参数
// 当前页码，每页几条数据
func GetArticleList(indexNum,pageSize int) (articleList []*model.ArticleInfo, err error) {
	// 从几开始indexNum取几条pageSize
	if indexNum < 0 || pageSize < 0{
		return
	}
	err = DB.Select(&articleList, "select id,summary,title,view_count,create_time,comment_count,username,category_id " +
		"from article where status=1 order by create_time desc limit ?,?", indexNum, pageSize)
	if err != nil{
		fmt.Println(err)
		return
	}
	return
}

// 根据文章id，查看文章信息，用于展示详情页
func GetArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error)  {
	articleDetail = &model.ArticleDetail{}
	err = DB.Get(articleDetail, "select id,summary,title,view_count,content,create_time,comment_count,username,category_id " +
		"from article where status=1 and id=?", articleId)
	if err != nil{
		fmt.Println(err)
		return
	}
	return
}

// 根据分类id，查文章
func GetArticleListByCategoryId(categoryId,indexNum,pageSize int64) (articleList []*model.ArticleInfo, err error)  {
	// 从几开始indexNum取几条pageSize
	if indexNum < 0 || pageSize < 0{
		return
	}
	err = DB.Select(&articleList, "select id,summary,title,view_count,create_time,comment_count,username,category_id " +
		"from article where status=1 and category_id=? order by create_time desc limit ?,?",categoryId, indexNum, pageSize)
	if err != nil{
		fmt.Println(err)
		return
	}
	return
}

// 获取相关文章
func GetRelativeArticle(articleId int64) (articleList []*model.RelativeArticle, err error) {
	var categoryId int64
	sqlstr := "select category_id from article where id=?"
	err = DB.Get(&categoryId, sqlstr, articleId)
	if err != nil {
		return
	}
	sqlstr = "select id, title from article where category_id=? and id !=?  limit 10"
	err = DB.Select(&articleList, sqlstr, categoryId, articleId)
	return
}

// 获取上一篇文章
func GetPrevArticleById(articleId int64) (info *model.RelativeArticle, err error)  {
	info = &model.RelativeArticle{
		ArticleId: -1,
	}
	sqlStr := "select id,title from article where id < ? order by id desc limit 1"
	err = DB.Get(info, sqlStr, articleId)
	if err != nil{
		fmt.Println(err)
		return
	}
	return
}

// 获取下一篇文章
func GetNextArticleById(articleId int64) (info *model.RelativeArticle, err error)  {
	info = &model.RelativeArticle{
		ArticleId: -1,
	}
	sqlStr := "select id,title from article where id > ? order by id asc limit 1"
	err = DB.Get(info, sqlStr, articleId)
	if err != nil{
		fmt.Println(err)
		return
	}
	return
}

// 判断文章是不是存在
func IsArticleExist(articleId int64) (exists bool, err error) {
	var id int64
	sqlstr := "select id from article where id=?"
	err = DB.Get(&id, sqlstr, articleId)
	if err == sql.ErrNoRows {
		exists = false
		return
	}
	if err != nil {
		return
	}
	exists = true
	return
}