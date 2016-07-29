package routers

import (
	"github.com/astaxie/beego"
	"github.com/dionyself/golang-cms/controllers"
)

//CTRL.db = utils.Mdb.Orm

func init() {
	//beego.DirectoryIndex = true
	// static routers
	//beego.ViewsPath = "views/default"
	//beego.SetViewsPath(utils.GetActiveView())
	beego.SetStaticPath("/static", "static/")
	//beego.SetStaticPath("/tpl/css", utils.GetActiveStylesPath())
	//beego.SetStaticPath("/down1", "download1")
	//beego.SetStaticPath("/down2", "download2")

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
