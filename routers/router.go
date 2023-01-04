package routers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/dionyself/golang-cms/controllers"
	"github.com/dionyself/golang-cms/core/template"
)

func init() {
	for template, styles := range template.Templates {
		for _, style := range styles {
			// web.BConfig.WebConfig.StaticDir
			web.SetStaticPath("/static/"+template+"/"+style, "views/"+template+"/styles/"+style)
		}
	}

	// guests request
	web.Router("/", &controllers.IndexController{})
	web.Router("/login", &controllers.LoginController{}, "get:LoginView;post:Login")
	web.Router("/logout", &controllers.LoginController{}, "get:Logout")
	web.Router("/register", &controllers.LoginController{}, "get:RegisterView;post:Register")
	web.Router("/article/:id:int/:action:string", &controllers.ArticleController{})

	// User requests
	web.Router("/ajax/image/:id:int", &controllers.AjaxController{}, "get:GetImageUploadStatus;post:PostImage")
	web.Router("/profile/:id:int/:action:string", &controllers.ProfileController{}, "get:UserPanelView")

	// filters
	web.InsertFilter("/profile/:id:int/show", web.BeforeRouter, controllers.AuthRequest)
	web.InsertFilter("/article/:id:int/edit", web.BeforeRouter, controllers.AuthRequest)
	web.InsertFilter("/article/:id:int/comment", web.BeforeRouter, controllers.AuthRequest)
	web.InsertFilter("/ajax/image/:id:int", web.BeforeRouter, controllers.AuthRequest)
	web.InsertFilter("/*", web.BeforeExec, controllers.DetectUserAgent)
	web.InsertFilter("/", web.BeforeExec, controllers.DetectUserAgent)
}
