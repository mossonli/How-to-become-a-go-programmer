package dao

import (
	"fmt"
	"goBlog/model"
	"testing"
	"time"
)

// 数据库初始化
func init() {
	err := Init("root:123456@tcp(127.0.0.1:3306)/blogger?parseTime=true")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestInsertArticle(t *testing.T) {
	// 构造文章对象
	article := &model.ArticleDetail{}
	article.ArticleInfo.CategoryId = 1
	article.ArticleInfo.CommentCount = 0
	article.Content = "这里是文章的内容"
	article.ArticleInfo.CreateTime = time.Now()
	article.ArticleInfo.Summary = "文章总结"
	article.ArticleInfo.Title = "文章标题"
	article.ArticleInfo.Username = "Mosson"
	article.ViewCount = 1
	articleId, err := InsertArticle(article)
	fmt.Println(articleId, err)
}

// 获取文章详细内容
func TestGetArticleDetail(t *testing.T) {
	articleDetail, _ := GetArticleDetail(3)
	fmt.Println(articleDetail)
}

func TestGetArticleListByCategoryIde(t *testing.T) {
	articleList, _ := GetArticleListByCategoryId(1, 0 ,10)
	fmt.Println(articleList)
}