package defaults

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/dionyself/golang-cms/core/defaults/modules"
)

// GetDefaultMenu get menu
func GetDefaultMenu() string {
	var menuitems string
	for _, mod := range modules.Modules {
		modConfig, err := web.AppConfig.GetSection("module-" + mod.Name)
		if err == nil && modConfig["activated"] != "" && modConfig["hidden"] != "" {
			menuitems = mod.Menu
		}
	}
	return menuitems
}
