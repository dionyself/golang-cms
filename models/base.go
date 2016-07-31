package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

// BaseForm ...
type BaseForm struct {
	InvalidFields map[string]string `form:"-"`
}

// Validate form data
func (form *BaseForm) Validate(data interface{}) bool {
	valid := validation.Validation{}
	b, err := valid.Valid(data)
	if err != nil {
		beego.Error(err)
	}
	if !b {
		if form.InvalidFields == nil {
			form.InvalidFields = make(map[string]string, len(valid.Errors))
		}
		for _, err := range valid.Errors {
			beego.Debug(err.Key, err.Message)
			form.InvalidFields[err.Key] = err.Message
		}
	}
	return b
}

func init() {
	//DB.Using("default")
}
