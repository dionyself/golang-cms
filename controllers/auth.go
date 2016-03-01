package controllers

import (
	"fmt"

	"github.com/dionyself/golang-cms/models"
	"github.com/dionyself/golang-cms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
)

var sessionName = beego.AppConfig.String("SessionName")

type LoginController struct {
	BaseController
}

func (this *LoginController) LoginView() {
	this.ConfigPage("login.html")

}

func (this *LoginController) Login() {
	username := this.GetString("username")
	password := this.GetString("password")
	backTo := this.GetString("back_to")

	var user models.User
	if VerifyUser(&user, username, password) {
		//session_data := this.GetSession(sessionName)
		this.SetSession(sessionName, user.Id)
		this.Redirect("/"+backTo, 302)

	} else {
		this.Redirect("/register", 302)
	}

}

func (this *LoginController) Logout() {
	this.DelSession(sessionName)
	this.Redirect("/login", 302)
}

func (this *LoginController) RegisterView() {
	this.ConfigPage("register.html")
}

func (this *LoginController) Register() {
	form := models.RegisterForm{}
	if err := this.ParseForm(&form); err != nil {
		this.Abort("401")
	}

	if form.Validate() {
		salt := utils.GetRandomString(10)
		encodedPwd := salt + "$" + utils.EncodePassword(form.Password, salt)

		o := this.GetDB()
		profile := new(models.Profile)
		profile.Age = 30
		user := new(models.User)
		user.Profile = profile
		user.Username = form.Username
		user.Password = encodedPwd
		user.Rands = salt
		fmt.Println(o.Insert(profile))
		fmt.Println(o.Insert(user))

		this.Redirect("/", 302)

	} else {
		this.ConfigPage("register.html")
	}
}

func HasUser(user *models.User, username string) bool {
	var err error
	qs := orm.NewOrm()
	user.Username = username
	err = qs.Read(user, "Username")
	if err == nil {
		return true
	}
	return false
}

func VerifyPassword(rawPwd, encodedPwd string) bool {
	var salt, encoded string
	salt = encodedPwd[:10]
	encoded = encodedPwd[11:]
	return utils.EncodePassword(rawPwd, salt) == encoded
}

func VerifyUser(user *models.User, username, password string) (success bool) {
	// search user by username or email
	if HasUser(user, username) == false {
		return
	}
	if VerifyPassword(password, user.Password) {
		success = true
	}
	return
}

var AuthRequest = func(ctx *context.Context) {
	uid, ok := ctx.Input.Session(sessionName).(int)
	if !ok && ctx.Input.URI() != "/login" && ctx.Input.URI() != "/register" {
		ctx.Redirect(302, "/login")
	}
	var user models.User
	var err error
	qs := orm.NewOrm()
	user.Id = uid
	err = qs.Read(&user, "Id")
	if err != nil {
		ctx.Redirect(302, "/login")
	}
	ctx.Input.SetData("user", user)
}
