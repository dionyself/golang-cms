package models

import (
	"github.com/astaxie/beego/validation"
)

type RegisterForm struct {
	BaseForm
	Name          string            `form:"name" valid:"Required;"`
	Email         string            `form:"email" valid:"Required;"`
	Username      string            `form:"username" valid:"Required;AlphaNumeric;MinSize(4);MaxSize(300)"`
	Password      string            `form:"password" valid:"Required;MinSize(4);MaxSize(30)"`
	PasswordRe    string            `form:"passwordre" valid:"Required;MinSize(4);MaxSize(30)"`
	Gender        bool              `form:"gender" valid:"Required"`
	InvalidFields map[string]string `form:"-"`
}

func (form *RegisterForm) Validate() bool {
	valid := form.BaseForm.Validate(form)
	if valid != true {
		form.InvalidFields = form.BaseForm.InvalidFields
	}
	return valid
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
	Id            int               `form:"-"`
	Title         string            `form:"title" valid:"Required;MinSize(4);MaxSize(300)"`
	Category      int               `form:"category"`
	Content       string            `form:"content" valid:"Required;MinSize(50);MaxSize(2000)"`
	TopicTags     string            `form:"topic-tags" valid:"MinSize(4);MaxSize(300)"`
	TaggedUsers   string            `form:"tagged-users" valid:"MinSize(4);MaxSize(300)"`
	AllowReviews  bool              `form:"allow-reviews" valid:"Required"`
	AllowComments bool              `form:"allow-comments" valid:"Required"`
	InvalidFields map[string]string `form:"-"`
}

func (form *ArticleForm) Validate() bool {
	valid := form.BaseForm.Validate(form)
	if valid != true {
		form.InvalidFields = form.BaseForm.InvalidFields
	}
	return valid
}

func (form *ArticleForm) Valid(v *validation.Validation) {
	if form.Category < 0 {
		v.SetError("Category", "Invalid category")
		return
	}
}
