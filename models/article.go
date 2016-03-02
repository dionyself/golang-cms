package models

import (
	"time"
)

type Article struct {
	Id             int       `orm:"column(id);auto"`
	User           *User     `orm:"rel(fk)"`
	Title          string    `orm:"column(title);size(255);"`
	Content        string    `orm:"column(content);size(128)"`
	CreateTime     time.Time `orm:"column(create_time);type(timestamp);auto_now_add"`
	Type           int
	Stars          int // we may need redis help with this
	AllowComments  bool
	Category       *Category         `orm:"rel(fk);null;default(null)"`
	ArticleComment []*ArticleComment `orm:"reverse(many)"`
	Likes          []*ArticleLike    `orm:"reverse(many)"`
}

type ArticleComment struct {
	Id      int            `orm:"column(id);auto"`
	User    *User          `orm:"rel(fk)"`
	Article *Article       `orm:"rel(fk)"`
	Likes   []*CommentLike `orm:"reverse(many)"`
}

type Category struct {
	Id       int        `orm:"column(id);auto"`
	Name     string     `orm:"column(name);size(128)"`
	Articles []*Article `orm:"reverse(many)"`
}
