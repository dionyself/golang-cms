package controllers

import (
	"github.com/Shaked/gomobiledetect"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dionyself/golang-cms/lib/defaults"
)

type MainController struct {
	beego.Controller
}

func (main *MainController) BeforeRender() {
	main.Layout = "layout.html"
	device := main.Ctx.Input.GetData("device_type").(string)
	main.LayoutSections = make(map[string]string)
	main.LayoutSections["Head"] = "partial/html_head_" + device + ".html"
	main.Data["menu_elements"] = main.GetMenu()
}

func (main *MainController) GetMenu() string {
	output := defaults.GetDefaultMenu()
	return output
}

func (main *MainController) GetContent() string {
	output := defaults.GetDefaultMenu()
	return output
}

func (index *MainController) Get() {
	index.Data["Website"] = "127.0.0.1:8080"
	index.Data["description"] = "Fastest and stable CMS"
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
