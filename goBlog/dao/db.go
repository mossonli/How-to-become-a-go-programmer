package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

// 数据库初始化
func Init(dns string) (err error) {
	DB, err = sqlx.Connect("mysql", dns)
	if err != nil{
		fmt.Println(err)
		return err
	}
	return nil
}