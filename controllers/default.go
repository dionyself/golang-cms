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
	//main_layout.LayoutSections["Scripts"] = "partial/scripts_" + main_layout.GetString("device_type") + ".html"
	//main_layout.LayoutSections["Styles"] = "partial/css_" + main_layout.GetString("device_type") + ".html"
	main.Data["menu_elements"] = main.GetMenu()
}

func (main *MainController) GetMenu() []map[string]string {
	output := defaults.GetDefaultMenu()
	return output
}

func (index *MainController) Get() {
	index.Data["Website"] = "127.0.0.1:8080"
	index.Data["description"] = "Fastest and stable CMS"
	index.Data["Email"] = "dionyself@gmail.com"
	index.TplNames = "index.html"
	index.BeforeRender()
}

var DetectUserAgent = func(ctx *context.Context) {
	deviceDetector := mobiledetect.NewMobileDetect(ctx.Request, nil)
	device, _ := ctx.Input.GetData("device_type").(string)
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
