package models

import (
	"time"
)

type Template struct {
	Id         int       `orm:"column(id);auto"`
	Name       string    `orm:"column(name);size(50);unique"`
	Style      []*Style  `orm:"reverse(many)"`
	Active     bool      `orm:"column(active)"`
	CreateTime time.Time `orm:"column(create_time);type(timestamp);auto_now_add"`
}

type Style struct {
	Id         int       `orm:"column(id);auto"`
	Name       string    `orm:"column(name);size(50);unique"`
	Template   *Template `orm:"rel(fk)"`
	Active     bool      `orm:"column(active)"`
	CreateTime time.Time `orm:"column(create_time);type(timestamp);auto_now_add"`
}
