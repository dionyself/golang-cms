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
	index.Data["Website"] = "Golang CMS, Fast CMS"
	index.Data["description"] = "Fastest and stable cms"
	index.Data["Email"] = "dionyself@gmail.com"
	index.TplNames = "index.tpl"
}

var DetectUserAgent = func(ctx *context.Context) {
	detector := mobiledetect.NewMobileDetect(ctx.Request, nil)
	session_data := ctx.Input.Session(sessionName)
	if session_data == nil {
		m := make(map[string]interface{})
		m["custom_theme"] = "default"
		m["custom_view"] = "mobile"
		m["custom_lang"] = "en"
	}
	if detector.IsMobile() {
		_, ok := ctx.Input.Session(sessionName).(int)
	}
	if detector.IsTablet() {
		_, ok := ctx.Input.Session(sessionName).(int)
	}
}
