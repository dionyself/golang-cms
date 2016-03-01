package models

import (
	"time"
)

type ArticleLike struct {
	Id         int       `orm:"column(id);auto"`
	User       *User     `orm:"rel(one)"`
	CreateTime time.Time `orm:"column(create_time);type(timestamp);auto_now_add"`
	Article    *Article  `orm:"rel(one)"`
}

type CommentLike struct {
	Id         int             `orm:"column(id);auto"`
	User       *User           `orm:"rel(one)"`
	CreateTime time.Time       `orm:"column(create_time);type(timestamp);auto_now_add"`
	Comment    *ArticleComment `orm:"rel(one)"`
}
