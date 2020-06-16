package model

import "time"

// 留言
type Leave struct {
	Id         int64     `db:"id"`
	Content    string    `db:"content"`
	Username   string    `db:"username"`
	CreateTime time.Time `db:"create_time"`
	Email      string    `db:"email"`
}
