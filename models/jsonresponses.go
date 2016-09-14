package models

// UploadResultJSON model articles in db
type UploadResultJSON struct {
	Id  int    `json:"id"`
	Msg string `json:"message"`
	Url string `orm:"column(url);size(255);"`
}
