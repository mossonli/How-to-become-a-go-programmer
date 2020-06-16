package service

import (
	"fmt"
	"goBlog/dao"
	"goBlog/model"
	"math"
)

// 获取所有文章 和对应的分类信息
func GetArticleRecordList(pageNum, pageSize int)(articleRecordList []*model.ArticleRecord, err error)  {
	// 1 获取所有的文章
	articleInfoList, err := dao.GetArticleList(pageNum, pageSize)
	if err != nil{
		return
	}
	// 2 获取文章对应的分类列表
	categoryIds := GetCategoryIds(articleInfoList)
	// 3 根据多个分类id 查询多个分类
	categoryList, err := dao.GetCategoryList(categoryIds)
	if err != nil{
		fmt.Println(err)
		return
	}
	// 4 聚合文章和对应的分类信息
	for _, article := range articleInfoList{
		articleRecord := &model.ArticleRecord{
			ArticleInfo:*article,
		}
		// 当前文章，取出分类id
		categoryId := article.CategoryId
		// 遍历分类id 对比数据
		for _,category := range categoryList{
			if categoryId == category.CategoryId{
				articleRecord.Category = *category
				break
			}
		}
		// 聚合之后的数据，加入返回值
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

//获取所有文章的分类信息
func GetCategoryIds(articleInfoList []*model.ArticleInfo) (ids []int64){
LABEL:
	for _, article := range articleInfoList{
		categoryId := article.CategoryId
		// 分类的id 不重复
		for _, id := range ids{
			if id == categoryId{
				continue LABEL
			}

		}
		ids = append(ids, categoryId)
	}
	return
}

// 根据分类id，获取该分类文章的列表和对应的分类信息
func GetArticleRecordListById(categoryId,pageNum,pageSize int64)(articleRecordList []*model.ArticleRecord, err error)  {
	// 1 根据分类id获取所有的文章
	articleInfoList, err := dao.GetArticleListByCategoryId(categoryId, pageNum, pageSize)
	if err != nil{
		fmt.Println(err)
		return
	}
	// 2 获取文章对应的分类列表
	categoryIds := GetCategoryIds(articleInfoList)
	// 3 根据多个分类的id，查询多个分类的信息
	categoryList, err := dao.GetCategoryList(categoryIds)
	if err != nil{
		fmt.Println(err)
		return
	}
	// 4 聚合文章和对应的分类信息
	for _, article := range articleInfoList{
		// 跟进文章，构建结构体
		articleRecord := &model.ArticleRecord{
			ArticleInfo:*article,
		}
		// 根据当前文章，取出分类id
		categoryId := article.CategoryId
		// 遍历分类id 对比数据
		for _,category := range categoryList{
			if categoryId == category.CategoryId{
				articleRecord.Category = *category
				break
			}
		}
		// 聚合之后的数据，加入返回值
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}
//
func GetRelativeArticleList(articleId int64) (articleList []*model.RelativeArticle, err error){
	articleList, err = dao.GetRelativeArticle(articleId)
	return
}
// 获取文章详情
func GetArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	articleDetail, err = dao.GetArticleDetail(articleId)
	if err != nil {
		return
	}
	category, err := dao.GetCategoryById(articleDetail.ArticleInfo.CategoryId)
	if err != nil {
		return
	}
	articleDetail.Category = *category
	return
}

// 获取文章的上一页 下一页
func GetPrevAndNextArticleInfo(articleId int64) (prevArticle *model.RelativeArticle, nextArticle *model.RelativeArticle, err error) {
	prevArticle, err = dao.GetPrevArticleById(articleId)
	if err != nil{
		fmt.Println(err)
		return
	}
	nextArticle, err = dao.GetNextArticleById(articleId)
	if err != nil{
		fmt.Println(err)
		return
	}
	return
}


// 文章保存
func InsertArticle(content,author,title string, categoryId int64) (err error) {
	articleDetail := &model.ArticleDetail{}
	// 组装对象
	articleDetail.Content = content
	articleDetail.Username = author
	articleDetail.Title = title
	articleDetail.ArticleInfo.CategoryId = categoryId
	// 从content切出摘要信息
	contentUtf8 := []rune(content)
	minLength := int(math.Min(float64(len(contentUtf8)), 128.0))
	articleDetail.Summary = string([]rune(content)[:minLength])
	id, err := dao.InsertArticle(articleDetail)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(id, err)
	return
}
