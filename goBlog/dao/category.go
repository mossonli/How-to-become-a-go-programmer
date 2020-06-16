package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"goBlog/model"
)

// 获取单个分类,返回一个category的对象
func GetCategoryById(id int64) (category *model.Category, err error)  {
	category = &model.Category{}
	DB.Get(category, "select id, category_name, category_no from category where id=?", id)
	return
}

// 获取多个分类
func GetCategoryList(categoryIds []int64)(categoryList []*model.Category, err error)  {
	// sqlx.In 构建可变参数去查询
	sqlStr, args, err := sqlx.In("select id, category_name, category_no from category where id IN(?)", categoryIds)
	err = DB.Select(&categoryList, sqlStr, args...)
	return
}

// 获取所有的分类
func GetAllCategoryList()(categoryList []*model.Category, err error){
	DB.Select(&categoryList, "select id, category_name, category_no from category")
	return
}