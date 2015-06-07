package routers

import (
	"github.com/astaxie/beego"
	"github.com/dionyself/golang-cms/controllers"
)

func init() {
	beego.DirectoryIndex = true
	// static routers
	beego.SetStaticPath("/static", "static/")

	// guests request
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{}, "get:LoginView;post:Login")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("/register", &controllers.LoginController{}, "get:RegisterView;post:Register")

	// User requests
	beego.Router("/my-account", &controllers.LoginController{}, "get:UserPanelView")

	// filters
	beego.InsertFilter("/my-account", beego.BeforeRouter, controllers.AuthRequest)
	beego.InsertFilter("/", beego.BeforeExec, controllers.DetectUserAgent)
}
