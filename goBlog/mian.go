package main

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"goBlog/controller"
	"goBlog/dao"
)

func main()  {
	router := gin.Default()
	dns := "root:123456@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	err := dao.Init(dns)
	if err != nil{
		panic(err)
	}
	// 集成性能测试工具
	ginpprof.Wrapper(router)
	// 加载静态资源
	router.Static("/static/", "./static")
	router.LoadHTMLGlob("views/*")
	//主页
	router.GET("/", controller.IndexHandler)
	// 分类
	router.GET("/category/", controller.CategoryList)
	// 点击投稿，去投稿页面
	router.GET("/article/new/", controller.NewArticle)
	// 添加投稿
	router.POST("/article/submit/", controller.ArticleSubmit)
	// 稿件详情
	router.GET("/article/detail/", controller.ArticleDetail)
	// 文件上传
	router.POST("/upload/file/", controller.UploadFile)
	// 新评论
	router.GET("/leave/new/", controller.LeaveNew)
	// 个人中心
	router.GET("/about/me/", controller.AboutMe)
	//
	router.POST("/comment/submit/", controller.CommentSubmit)
	router.POST("/leave/submit/", controller.LeaveSubmit)
	_ = router.Run(":8000")
}