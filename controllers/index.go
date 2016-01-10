package controllers

import (

)


type IndexController struct {
	BaseController
}

func (index *IndexController) Get() {
	index.ConfigPage("index.html")
	index.Data["Website"] = "127.0.0.1:8080"
	index.Data["description"] = "Fast and stable CMS"
	// index.Data["content"] = index.getContent()
	index.Data["Email"] = "dionyself@gmail.com"
}