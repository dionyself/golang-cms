package controllers

type IndexController struct {
	BaseController
}

// Get main page
func (CTRL *IndexController) Get() {
	CTRL.ConfigPage("index.html")
	CTRL.Data["Website"] = "127.0.0.1:8080"
	CTRL.Data["description"] = "Fast and stable CMS"
	CTRL.Data["Email"] = "dionyself@gmail.com"
}
