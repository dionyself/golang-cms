package modules

var Modules map[string]map[string]string

func init() {
	Modules = make(map[string]map[string]string)
	news := make(map[string]string)
	news["menu"] = `{"news": "/", "extra": "/"}`
	Modules["news"] = news
}
