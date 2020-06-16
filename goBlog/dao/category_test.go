package dao

import (
	"fmt"
	"testing"
)

// 数据库初始化
func init()  {
	err := Init("root:123456@tcp(127.0.0.1:3306)/blogger?parseTime=true")
	if err != nil{
		fmt.Println(err)
		return
	}
}

// 单元测试 查单个
func TestGetCategoryById(t *testing.T) {
	category, err := GetCategoryById(1)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(category)
}

// 单元测试 查多个
func TestGetCategoryList(t *testing.T) {
	var categoryIds []int64
	categoryIds = append(categoryIds, 1, 2, 3)
	categoryList, err := GetCategoryList(categoryIds)
	if err != nil{
		fmt.Println(err)
		return
	}
	for _, v := range categoryList{
		fmt.Printf("id:%d category %#v\n", v.CategoryId, v)
	}
}

// 查询所有
func TestGetAllCategoryList(t *testing.T) {
	categoryList, err := GetAllCategoryList()
	if err != nil{
		fmt.Println(err)
		return
	}
	for _, v := range categoryList{
		fmt.Printf("id:%d category %#v\n", v.CategoryId, v)
	}
}