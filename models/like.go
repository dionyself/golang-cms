package models

import (
	"time"
)

// ArticleLike model
type ArticleLike struct {
	Id         int       `orm:"column(id);auto"`
	User       *User     `orm:"rel(fk)"`
	CreateTime time.Time `orm:"column(create_time);type(timestamp);auto_now_add"`
	Article    *Article  `orm:"rel(fk)"`
}

// CommentLike model
type CommentLike struct {
	Id         int             `orm:"column(id);auto"`
	User       *User           `orm:"rel(fk)"`
	CreateTime time.Time       `orm:"column(create_time);type(timestamp);auto_now_add"`
	Comment    *ArticleComment `orm:"rel(fk)"`
}
