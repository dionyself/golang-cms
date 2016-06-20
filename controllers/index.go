package controllers

import "github.com/dionyself/golang-cms/utils"

type IndexController struct {
	BaseController
}

func (CTRL *IndexController) Get() {
	CTRL.ConfigPage("index.html")
	utils.Mcache.Set("test", "127.0.0.1:8080", 60)
	ip, err := utils.Mcache.GetString("test", 60)
	if err != false {
		CTRL.Data["Website"] = ip + "test"
	}
	CTRL.Data["description"] = "Fast and stable CMS"
	// CTRL.Data["content"] = CTRL.getContent()
	CTRL.Data["Email"] = "dionyself@gmail.com"
}
