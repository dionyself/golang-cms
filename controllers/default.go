package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "Golang CMS, Fast CMS"
	this.Data["description"] = "Neutral reviews, check if product is good, find quality products, compare products"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}
