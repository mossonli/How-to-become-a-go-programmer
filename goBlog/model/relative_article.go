package model
// 文章的上一页 下一页
type RelativeArticle struct {
	ArticleId int64  `db:"id"`
	Title     string `db:"title"`
}
