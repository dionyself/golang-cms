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
	beego.Router("/offers", &controllers.MainController{})

	// User requests
	beego.Router("/secret", &controllers.LoginController{}, "get:UserPanelView")
	beego.Router("/my-account", &controllers.LoginController{}, "get:UserPanelView")
	beego.Router("/user-offers", &controllers.MainController{})

	// Vendor requests
	beego.Router("/vendor-panel", &controllers.VendorPanelController{})
	beego.Router("/vendor-offers", &controllers.MainController{})

	// admin requests
	beego.Router("/admin-panel", &controllers.AdminPanelController{})
	beego.Router("/admin-offers", &controllers.MainController{})

	// filters
	beego.InsertFilter("/vendor-panel", beego.BeforeRouter, controllers.AuthRequest)
	beego.InsertFilter("/user-panel", beego.BeforeRouter, controllers.AuthRequest)
	beego.InsertFilter("/my-account", beego.BeforeRouter, controllers.AuthRequest)
	beego.InsertFilter("/", beego.BeforeExec, controllers.DetectUserAgent)
}
