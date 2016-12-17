package models

import (
	"time"
)

// Block model blocks in db
type Block struct {
	Id         int       `orm:"column(id);auto"`
	Name       string    `orm:"column(title);size(255);"`
	Content    string    `orm:"column(content);size(128)"`
	CreateTime time.Time `orm:"column(create_time);type(timestamp);auto_now_add"`
	Type       string    `orm:"column(type);size(128)"`
	IsActive   bool
	Position   int
	//Config     []*BlockConfig `orm:"reverse(many)"`
}

// BlockConfig model
type BlockConfig struct {
	Id         int            `orm:"column(id);auto"`
	Block      *Block         `orm:"rel(fk)"`
	Key        string         `orm:"column(title);size(255);"`
	Value      string         `orm:"column(title);size(255);"`
	Type       string         `orm:"column(title);size(255);"`
	CreateTime time.Time      `orm:"column(create_time);type(timestamp);auto_now_add"`
	Config     []*BlockConfig `orm:"reverse(many)"`
}
