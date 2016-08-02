package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

// RegisterForm ...
type RegisterForm struct {
	Name          string            `form:"name" valid:"Required;"`
	Email         string            `form:"email" valid:"Required;"`
	Username      string            `form:"username" valid:"Required;AlphaNumeric;MinSize(4);MaxSize(300)"`
	Password      string            `form:"password" valid:"Required;MinSize(4);MaxSize(30)"`
	PasswordRe    string            `form:"passwordre" valid:"Required;MinSize(4);MaxSize(30)"`
	Gender        bool              `form:"gender" valid:"Required"`
	InvalidFields map[string]string `form:"-"`
}

// Validate form data
func (form *RegisterForm) Validate() bool {
	validator := validation.Validation{}
	isValid := false
	var err error
	if isValid, err = validator.Valid(form); err != nil {
		beego.Error(err)
	} else {
		if !isValid {
			if form.InvalidFields == nil {
				form.InvalidFields = make(map[string]string, len(validator.Errors))
			}
			for _, err := range validator.Errors {
				beego.Debug(err.Key, err.Message)
				form.InvalidFields[err.Key] = err.Message
			}
		}
	}
	return isValid
}

// Valid check if RegisterForm is valid
func (form *RegisterForm) Valid(v *validation.Validation) {
	// Check if passwords of two times are same.
	if form.Password != form.PasswordRe {
		v.SetError("PasswordRe", "Passwords did not match")
		return
	}
}

// ArticleForm ...
type ArticleForm struct {
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

// Validate form data
func (form *ArticleForm) Validate() bool {
	validator := validation.Validation{}
	isValid := false
	var err error
	if isValid, err = validator.Valid(form); err != nil {
		beego.Error(err)
	} else {
		if !isValid {
			if form.InvalidFields == nil {
				form.InvalidFields = make(map[string]string, len(validator.Errors))
			}
			for _, err := range validator.Errors {
				beego.Debug(err.Key, err.Message)
				form.InvalidFields[err.Key] = err.Message
			}
		}
	}
	return isValid
}

// Valid checks if ArticleForm is valid
func (form *ArticleForm) Valid(v *validation.Validation) {
	if form.Category < 0 {
		v.SetError("Category", "Invalid category")
		return
	}
}
