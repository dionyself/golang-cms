package controllers

import (
	"github.com/astaxie/beego"

	/*
		"fmt"
		"github.com/astaxie/beego/context"
		"github.com/astaxie/beego/orm"
		"github.com/astaxie/beego/validation"
		_ "github.com/go-sql-driver/mysql"
		"github.com/dionyself/golang-cms/models"
		"github.com/dionyself/golang-cms/utils"
	*/)

type UserPanelController struct {
	beego.Controller
}

func (this *UserPanelController) MainView() {
	this.Layout = "layout.html"
	this.TplNames = "user-panel.html"
}

type VendorPanelController struct {
	beego.Controller
}
type AdminPanelController struct {
	beego.Controller
}
