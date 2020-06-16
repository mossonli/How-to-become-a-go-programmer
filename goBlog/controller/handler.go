package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"goBlog/service"
	"goBlog/utils"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
)

// 主要是处理回调函数

// 访问主页的回调函数
func IndexHandler(c *gin.Context)  {
	articleRecordList, err := service.GetArticleRecordList(0, 10)
	if err != nil{
		c.HTML(http.StatusInternalServerError, "views/500.html", err )
		return
	}
	// 获取分类，用于分类云展示
	categoryList, err := service.GetAllCategoryList()
	if err != nil{
		c.HTML(http.StatusInternalServerError, "views/500.html", err)
	}
	c.HTML(200, "views/index.html", gin.H{
		"article_list": articleRecordList,
		"category_list": categoryList,
	})
}

// 分类云的回调函数
func CategoryList(c *gin.Context)  {
	// 点击分类按钮的id
	categoryIdStr := c.Query("category_id")
	categoryId, err := strconv.ParseInt(categoryIdStr,10, 64)
	if err != nil{
		c.HTML(http.StatusInternalServerError, "views/500.html", err)
		return
	}
	// 根据分类的id 查文章
	articleRecordList, err := service.GetArticleRecordListById(categoryId, 0, 10)
	if err != nil{
		c.HTML(http.StatusInternalServerError, "views/500.html", err)
		return
	}
	// 获取分类，用于分类云展示
	categoryList, err := service.GetAllCategoryList()
	if err != nil{
		c.HTML(http.StatusInternalServerError, "views/500.html", err)
	}
	c.HTML(200, "views/index.html", gin.H{
		"article_list": articleRecordList,
		"category_list": categoryList,
	})
}

// 点击投稿 跳转到投稿页面
func NewArticle(c *gin.Context)  {
	// 获取所有的分类信息
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		fmt.Printf("get article failed, err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	c.HTML(http.StatusOK, "views/post_article.html", categoryList)
}

// 插入文章, 完成投稿
func ArticleSubmit(c *gin.Context)  {
	// 从页面接收表单数据，调用service进行存储
	content := c.PostForm("content")
	author := c.PostForm("author")
	categoryIdStr := c.PostForm("category_id")
	title := c.PostForm("title")
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	// 存完之后，重定向到首页
	err = service.InsertArticle(content, author, title, categoryId)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	// 转发到首页
	c.Redirect(http.StatusMovedPermanently, "/")
}

// 详情页
func ArticleDetail(c *gin.Context)  {
	articleIdStr := c.Query("article_id")
	articleId, err := strconv.ParseInt(articleIdStr, 10, 64)
	if err != nil{
		c.HTML(http.StatusInternalServerError, "views/500.html", err)
		return
	}
	// 查出具体文章
	articleDetail, err := service.GetArticleDetail(articleId)
	if err != nil{
		c.HTML(http.StatusInternalServerError, "views/500.html", err)
		return
	}
	// 组装文章的详情页需要显示的信息
	relativeArticle, err := service.GetRelativeArticleList(articleId)
	if err != nil{
		fmt.Println(err)
	}
	// 上一篇和下一篇
	prevArticle, nextArticle, err := service.GetPrevAndNextArticleInfo(articleId)
	if err != nil{
		fmt.Println(err)
	}
	// 获取所有的分类
	allCategoryList, err := service.GetAllCategoryList()
	if err != nil{
		fmt.Println(err)
	}
	// 获取评论列表
	commentList, err := service.GetCommentList(articleId)
	if err != nil{
		fmt.Println(err)
	}
	var m map[string]interface{} = make(map[string]interface{}, 10)
	m["detail"] = articleDetail
	m["relative_article"] = relativeArticle
	m["prev"] = prevArticle
	m["next"] = nextArticle
	m["category"] = allCategoryList
	m["article_id"] = articleId
	m["comment_list"] = commentList
	c.HTML(http.StatusOK, "views/detail.html", m)
	return
}

// 上传文件
func UploadFile(c *gin.Context)  {
	// single file
	file, err := c.FormFile("upload")
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":err.Error(),
		})
		return
	}
	log.Println(file.Filename)
	rootPath := utils.GetRootDir()
	//u2 := uuid.NewV4()
	u2 := uuid.NewV4()
	//if err != nil {
	//	return
	//}
	ext := path.Ext(file.Filename)
	url := fmt.Sprintf("/static/upload/%s%s", u2, ext)
	dst := filepath.Join(rootPath, url)
	// Upload the file to specific dst.
	_ = c.SaveUploadedFile(file, dst)
	c.JSON(http.StatusOK, gin.H{
		"uploaded": true,
		"url":      url,
	})
}

// 提交评论
func CommentSubmit(c *gin.Context) {
	comment := c.PostForm("comment")
	author := c.PostForm("author")
	email := c.PostForm("email")
	articleIdStr := c.PostForm("article_id")
	articleId, err := strconv.ParseInt(articleIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	err = service.InsertComment(comment, author, email, articleId)
	if err != nil {
		fmt.Printf("insert comment failed, err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	url := fmt.Sprintf("/article/detail/?article_id=%d", articleId)
	c.Redirect(http.StatusMovedPermanently, url)
}

// 添加留言
func LeaveSubmit(c *gin.Context) {
	comment := c.PostForm("comment")
	author := c.PostForm("author")
	email := c.PostForm("email")
	err := service.InsertLeave(author, email, comment)
	if err != nil {
		fmt.Printf("insert leave failed, err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	url := fmt.Sprintf("/leave/new/")
	c.Redirect(http.StatusMovedPermanently, url)
}

// 新
func LeaveNew(c *gin.Context) {
	leaveList, err := service.GetLeaveList()
	if err != nil {
		fmt.Printf("get leave failed, err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/gbook.html", leaveList)
}

// 关于我
func AboutMe(c *gin.Context) {
	c.HTML(http.StatusOK, "views/about.html", gin.H{
		"title": "Posts",
	})
}