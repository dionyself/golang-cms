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
	device := ctx.Input.Cookie("Device-Type")
	if device == "" {
		device, ok := ctx.Input.GetData("device_type").(string)
		if ok {
			ctx.Output.Cookie("Device-Type", device)
		} else {
			if deviceDetector.IsMobile() {
				device = "Mobile"
			}
			if deviceDetector.IsTablet() {
				device = "Tablet"
			}
			if device != "" {
				ctx.Output.Cookie("Device-Type", device)
			} else {
				device = beego.AppConfig.String("DefaultVersion")
				ctx.Output.Cookie("Device-Type", device)
			}
		}
	}
	ctx.Input.SetData("device_type", device)
}
