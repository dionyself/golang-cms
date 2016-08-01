package routers

import (
	"github.com/astaxie/beego"
	"github.com/dionyself/golang-cms/controllers"
	"github.com/dionyself/golang-cms/utils"
)

func init() {
	for template, styles := range utils.Templates {
		for _, style := range styles {
			// beego.BConfig.WebConfig.StaticDir
			beego.SetStaticPath("/static/"+template+"/"+style, "views/"+template+"/styles/"+style)
		}
	}

	// guests request
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/login", &controllers.LoginController{}, "get:LoginView;post:Login")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("/register", &controllers.LoginController{}, "get:RegisterView;post:Register")
	beego.Router("/article/:id:int/:action:string", &controllers.ArticleController{})

	// User requests
	beego.Router("/profile/:id:int/:action:string", &controllers.ProfileController{}, "get:UserPanelView")

	// filters
	beego.InsertFilter("/profile/:id:int/show", beego.BeforeRouter, controllers.AuthRequest)
	beego.InsertFilter("/article/:id:int/edit", beego.BeforeRouter, controllers.AuthRequest)
	beego.InsertFilter("/article/:id:int/comment", beego.BeforeRouter, controllers.AuthRequest)
	beego.InsertFilter("/*", beego.BeforeExec, controllers.DetectUserAgent)
	beego.InsertFilter("/", beego.BeforeExec, controllers.DetectUserAgent)
}
