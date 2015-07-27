package defaults

import (
	"github.com/astaxie/beego"
	"github.com/dionyself/golang-cms/lib/defaults/modules" // refs
)

// GetDefaultMenu get menu
func GetDefaultMenu() []string {
	menuitems := []string{}
	for modulename, val := range modules.Modules {
		modConfig, err := beego.AppConfig.GetSection("module-" + modulename)
		if err == nil && modConfig["activated"] != "" && modConfig["hidden"] != "" {
			menuitems = append(menuitems, val["menu"])
		}
	}
	return menuitems
}
