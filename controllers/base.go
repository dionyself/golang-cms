package controllers

import (
	"strconv"

	mobiledetect "github.com/Shaked/gomobiledetect"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/dionyself/golang-cms/core/block"
	"github.com/dionyself/golang-cms/core/defaults"
	"github.com/dionyself/golang-cms/core/lib/cache"
	database "github.com/dionyself/golang-cms/core/lib/db"
	"github.com/dionyself/golang-cms/core/template"
	"github.com/dionyself/golang-cms/utils"
)

// BaseController Extendable
type BaseController struct {
	web.Controller
	db    orm.Ormer
	cache cache.CACHE
}

// ConfigPage receives template name and makes basic config to render it
func (CTRL *BaseController) ConfigPage(page string) {
	CTRL.GetDB()
	CTRL.GetCache()
	theme := template.GetActiveTheme(false)
	CTRL.Layout = theme[0] + "/" + "layout.html"
	device := CTRL.Ctx.Input.GetData("device_type").(string)
	CTRL.LayoutSections = make(map[string]string)
	CTRL.LayoutSections["Head"] = theme[0] + "/" + "partial/html_head_" + device + ".html"
	CTRL.TplName = theme[0] + "/" + page
	CTRL.Data["Theme"] = theme[0]
	CTRL.Data["Style"] = theme[1]
	CTRL.Data["ModuleMenu"] = CTRL.GetModuleMenu()
	CTRL.GetBlocks()
	//CTRL.GetActiveModule()
	//CTRL.GetActiveCategory()
	//CTRL.GetActiveAds()
}

func (CTRL *BaseController) GetBlocks() map[string]string {
	// TODO : get blocks and set block content
	// loadedBlocks := CTRL.cache.GetMap(fmt.Sprintf("activeblocks/%s/%s", module, session_id) , expirationTime)
	loadedBlocks := make(map[string]string)
	ActiveBlocks := block.GetActiveBlocks(false)
	for _, CurentBlock := range ActiveBlocks {
		cblock := block.Blocks[CurentBlock]
		cblockSectionName := "Block_" + strconv.Itoa(cblock.GetPosition())
		cblockSectionData := cblockSectionName + "_Data"
		loadedBlocks[cblockSectionName] = cblock.GetTemplatePath()
		CTRL.Data[cblockSectionData] = cblock.GetContent()
	}
	CTRL.LayoutSections = utils.MergeMaps(CTRL.LayoutSections, loadedBlocks)
	return CTRL.LayoutSections
}

func (CTRL *BaseController) GetActiveAds() map[string]string {
	// loadedAds := CTRL.cache.GetMap(fmt.Sprintf("activeAds/%s/%s", category, session_id) , expirationTime)
	return make(map[string]string)
}

// GetCache set the cache connector into our controller
func (CTRL *BaseController) GetCache() {
	CTRL.cache = cache.MainCache
}

// GetDB set the orm connector into our controller
// if repication activated we use slave to Slave
func (CTRL *BaseController) GetDB(db ...string) orm.Ormer {
	if len(db) > 0 {
		CTRL.db = database.MainDatabase.GetOrm(db[0])
	} else {
		CTRL.db = database.MainDatabase.GetOrm("")
	}
	return CTRL.db
}

// GetModuleMenu retrieves menu
func (CTRL *BaseController) GetModuleMenu() string {
	output := defaults.GetDefaultMenu()
	return output
}

// GetContent gets contents
func (CTRL *BaseController) GetContent() string {
	output := defaults.GetDefaultMenu()
	return output
}

// DetectUserAgent detects device type and set it into a cookie
var DetectUserAgent = func(ctx *context.Context) {
	deviceDetector := mobiledetect.NewMobileDetect(ctx.Request, nil)
	ctx.Request.ParseForm()
	device := ""
	if len(ctx.Request.Form["device_type"]) != 0 {
		device = ctx.Request.Form["device_type"][0]
	}
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
			device, _ = web.AppConfig.String("DefaultDevice")
			if device == "" {
				device = "desktop"
			}
		}
	}
	ctx.Output.Cookie("Device-Type", device)
	ctx.Input.SetData("device_type", device)
}
