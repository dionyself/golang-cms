package models

import (
	"time"
)

type Article struct {
	Id         int       `orm:"column(id);auto"`
	User       *User     `orm:"rel(one)"`
	Title      string    `orm:"column(title);size(255);"`
	Content    string    `orm:"column(content);size(128)"`
	CreateTime time.Time `orm:"column(create_time);type(timestamp);auto_now_add"`
	Type       int
	Category   *Category     `orm:"rel(one)"`
	ArticleComment *ArticleComment `orm:"reverse(one)"`
	
}

type ArticleComment struct {
	Id         int       `orm:"column(id);auto"`
	User       *User     `orm:"rel(one)"`
	Article    *Article  `orm:"rel(one)"`
	
}

type Category struct {
	Id         int       `orm:"column(id);auto"`
	Name       string    `orm:"column(namd);size(128)"`
	Article    *Article `orm:"reverse(one)"`
}
