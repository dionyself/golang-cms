package controllers

import (
	"github.com/Shaked/gomobiledetect"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type MainController struct {
	beego.Controller
}

func (index *MainController) Get() {
	index.Data["Website"] = "Golang-CMS, Fast CMS"
	index.Data["description"] = "Fastest and stable CMS"
	index.Data["Email"] = "dionyself@gmail.com"
	index.TplNames = "index.tpl"
}

var DetectUserAgent = func(ctx *context.Context) {
	deviceDetector := mobiledetect.NewMobileDetect(ctx.Request, nil)
	device, _ := ctx.Input.GetData("device_type").(string)
	if device == "" {
		device = ctx.Input.Cookie("Device-Type")
	}
	if device == "" {
		if deviceDetector.IsMobile() {
			device = "Mobile"
		}
		if deviceDetector.IsTablet() {
			device = "Tablet"
		}
		if device == "" {
			device = beego.AppConfig.String("DefaultDevice")
		}
	}
	ctx.Output.Cookie("Device-Type", device)
	ctx.Input.SetData("device_type", device)
}
