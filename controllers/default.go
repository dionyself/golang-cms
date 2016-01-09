package controllers

import (
	"github.com/dionyself/gomobiledetect"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dionyself/golang-cms/lib/defaults"
)

type BaseController struct {
	beego.Controller
}

type IndexController struct {
	BaseController
}

type ArticleController struct {
	BaseController
}

func (this *ArticleController) Get() {
	this.TplNames = "article-editor.html"
	this.BeforeRender()
}

func (base *BaseController) BeforeRender() {
	base.Layout = "layout.html"
	device := base.Ctx.Input.GetData("device_type").(string)
	base.LayoutSections = make(map[string]string)
	base.LayoutSections["Head"] = "partial/html_head_" + device + ".html"
	base.Data["menu_elements"] = base.GetMenu()
}


func (base *BaseController) GetMenu() string {
	output := defaults.GetDefaultMenu()
	return output
}

func (base *BaseController) GetContent() string {
	output := defaults.GetDefaultMenu()
	return output
}

func (index *IndexController) Get() {
	index.Data["Website"] = "127.0.0.1:8080"
	index.Data["description"] = "Fast and stable CMS"
	// index.Data["content"] = index.getContent()
	index.Data["Email"] = "dionyself@gmail.com"
	index.TplNames = "index.html"
	index.BeforeRender()
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
