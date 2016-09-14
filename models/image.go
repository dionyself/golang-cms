package models

// Image model articles in db
type Image struct {
	Id       int    `orm:"column(id);auto"`
	User     *User  `orm:"rel(fk)"`
	Title    string `orm:"column(title);size(255);"`
	Url      string `orm:"column(url);size(255);"`
	Type     int
	Category *Category `orm:"rel(fk);null;default(null)"`
}
