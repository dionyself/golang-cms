package controllers

import (
	//"github.com/astaxie/beego"

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
	BaseController
}

func (this *UserPanelController) MainView() {
	this.ConfigPage("user-panel.html")
}

type VendorPanelController struct {
	BaseController
}
type AdminPanelController struct {
	BaseController
}
