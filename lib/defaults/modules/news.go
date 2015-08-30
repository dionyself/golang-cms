package modules

/*
type ModuleConfig struct {
  Name string
  Menu int
  weight int
}
var ModulesConfig = []ModuleConfig
*/

var Modules map[string]map[string]map[string]string

func init() {
	Modules = make(map[string]map[string]map[string]string)
	news := make(map[string]map[string]string)
	news["menu"] = make(map[string]string)
	news["menu"]["news"] = "/"
	news["menu"]["extra"] = "/"
	Modules["news"] = news
}
