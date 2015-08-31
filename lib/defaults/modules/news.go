package modules

type ModuleConfig struct {
	Name   string
	Menu   string
	weight int
}

var Modules []ModuleConfig

func init() {
	var moduleConfig ModuleConfig
	moduleConfig.Name = "news"
	moduleConfig.Menu = "{'news': '/news', 'top_10': '/news_top'}"
	Modules = append(Modules, moduleConfig)
}
