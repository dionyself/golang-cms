package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type BaseForm struct {
	Errors map[string]string
}

func (form *BaseForm) Validate() bool {
	valid := validation.Validation{}
	b, err := valid.Valid(form)
	if err != nil {
		beego.Error(err)
	}
	if !b {
		for _, err := range valid.Errors {
			form.Errors[err.Key] = err.Message
			beego.Debug(err.Key, err.Message)
		}
	}
	return b
}

type RegisterForm struct {
	BaseForm
	Username   string `form:"username" valid:"Required; AlphaNumeric; MinSize(4); MaxSize(300)"`
	Password   string `form:"password" valid:"Required; MinSize(4); MaxSize(30)"`
	PasswordRe string `form:"passwordre" valid:"Required; MinSize(4); MaxSize(30)"`
}

func (form *RegisterForm) Valid(v *validation.Validation) {
	// Check if passwords of two times are same.
	if form.Password != form.PasswordRe {
		v.SetError("PasswordRe", "Passwords did not match")
		return
	}
}

type ArticleForm struct {
	BaseForm
	Id            int    `form:"-"`
	Title         string `form:"title" valid:"Required;MinSize(4);MaxSize(300)"`
	Category      int    `form:"category"`
	Content       string `form:"content" valid:"Required; MinSize(50); MaxSize(2000)"`
	TopicTags     string `form:"topic-tags" valid:"MinSize(4); MaxSize(300)"`
	TaggedUsers   string `form:"tagged-users" valid:"MinSize(4); MaxSize(300)"`
	AllowReviews  bool   `form:"allow-reviews" valid:"Required"`
	AllowComments bool   `form:"allow-comments" valid:"Required"`
	Errors        map[string]string
}

func (form *ArticleForm) Valid(v *validation.Validation) {
	if form.Category >= 0 {
		v.SetError("Category", "Invalid category")
		return
	}
}
