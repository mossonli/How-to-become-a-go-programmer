package service

import (
	"goBlog/dao"
	"goBlog/model"
)

func GetAllCategoryList() (categoryList []*model.Category, err error)  {
	// 调用dao层 获取数据
	categoryList, err = dao.GetAllCategoryList()
	if err != nil{
		return
	}
	return
}
