package models

import (
	"time"

	"github.com/dionyself/beego/orm"
)

// User model
type User struct {
	Id          int        `orm:"column(id);auto"`
	Username    string     `orm:"column(username);size(50);unique"`
	Email       string     `orm:"column(email);size(255);"`
	Password    string     `orm:"column(password);size(128)"`
	CreateTime  time.Time  `orm:"column(create_time);type(timestamp);auto_now_add"`
	Admin       bool       `orm:"column(admin)"`
	Rands       string     `orm:"size(10)"`
	Profile     *Profile   `orm:"rel(one)"`
	Article     []*Article `orm:"reverse(many)"`
	Permissions string
}

// Profile model
type Profile struct {
	Id          int
	User        *User `orm:"reverse(one)"`
	Name        string
	Avatar      string
	Age         int16
	Lema        string
	Description string
	Gender      bool
	//Socials   []*Social `orm:"reverse(many)"`
}

// GetPermissions get user permissions data
func (*Profile) GetPermissions(user User) string {
	return user.Permissions
}

// registering all modules
func init() {
	orm.RegisterModel(
		new(User),
		new(Profile),
		new(Article),
		new(ArticleComment),
		new(Category),
		new(ArticleLike),
		new(CommentLike),
		new(Template),
		new(Style),
		new(Image),
		new(Block),
	)
}
