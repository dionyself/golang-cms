package defaults

import (
	"github.com/dionyself/beego"
	"github.com/dionyself/golang-cms/core/defaults/modules"
)

// GetDefaultMenu get menu
func GetDefaultMenu() string {
	var menuitems string
	for _, mod := range modules.Modules {
		modConfig, err := beego.AppConfig.GetSection("module-" + mod.Name)
		if err == nil && modConfig["activated"] != "" && modConfig["hidden"] != "" {
			menuitems = mod.Menu
		}
	}
	return menuitems
}
