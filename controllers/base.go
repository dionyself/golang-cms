package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/dionyself/golang-cms/lib/defaults"
	"github.com/dionyself/golang-cms/utils"
	"github.com/dionyself/gomobiledetect"
)

type BaseController struct {
	beego.Controller
	db orm.Ormer
}

func (CTRL *BaseController) ConfigPage(page string) {
	theme := utils.GetActiveTheme(false)
	CTRL.Layout = theme[0] + "/" + "layout.html"
	device := CTRL.Ctx.Input.GetData("device_type").(string)
	CTRL.LayoutSections = make(map[string]string)
	CTRL.LayoutSections["Head"] = theme[0] + "/" + "partial/html_head_" + device + ".html"
	CTRL.TplName = theme[0] + "/" + page
	CTRL.Data["Style"] = theme[1]
	_ = CTRL.GetDB()
	// CTRL.Data["Title"] = CTRL.GetPageTitle()
	CTRL.Data["ModuleMenu"] = CTRL.GetModuleMenu()
}

func (CTRL *BaseController) GetDB(db ...string) orm.Ormer {
	CTRL.db = utils.Mdb.Orm
	if len(db) > 0 {
		CTRL.db.Using(db[0])
	}
	return CTRL.db
}

func (CTRL *BaseController) GetModuleMenu() string {
	output := defaults.GetDefaultMenu()
	return output
}

func (CTRL *BaseController) GetContent() string {
	output := defaults.GetDefaultMenu()
	return output
}

var DetectUserAgent = func(ctx *context.Context) {
	deviceDetector := mobiledetect.NewMobileDetect(ctx.Request, nil)
	ctx.Request.ParseForm()
	device := ""
	if len(ctx.Request.Form["device_type"]) != 0 {
		device = ctx.Request.Form["device_type"][0]
	}
	if device == "" {
		device = ctx.Input.Cookie("Device-Type")
	}
	if device == "" {
		if deviceDetector.IsMobile() {
			device = "mobile"
		}
		if deviceDetector.IsTablet() {
			device = "tablet"
		}
		if device == "" {
			device = beego.AppConfig.String("DefaultDevice")
			if device == "" {
				device = "desktop"
			}
		}
	}
	ctx.Output.Cookie("Device-Type", device)
	ctx.Input.SetData("device_type", device)
}
