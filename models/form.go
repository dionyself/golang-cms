package models

import (
	"github.com/astaxie/beego/validation"
)

type RegisterForm struct {
	Username   string `valid:"AlphaNumeric" valid:"Required;MinSize(4);MaxSize(300)"`
	Password   string `form:"type(password)" valid:"Required;MinSize(4);MaxSize(30)"`
	PasswordRe string `form:"type(password)" valid:"Required;MinSize(4);MaxSize(30)"`
}

func (form *RegisterForm) Valid(v *validation.Validation) {
	// Check if passwords of two times are same.
	if form.Password != form.PasswordRe {
		v.SetError("PasswordRe", "Passwords did not match")
		return
	}
}
