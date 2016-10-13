package block

import (
	_ "github.com/astaxie/beego"
	_ "github.com/astaxie/beego/orm"
	"github.com/dionyself/golang-cms/core/lib/cache"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// Blocks map["block_folder"]["style1", "style2" ...]
var Blocks map[string]Block
var RegisteredBlocks map[string]Block

type Block interface {
	GetTemplatePath() string
	GetContent() string
	GetPosition() string
	//IsContentCacheable() bool
	//GetConfig() models.BlockConfig
	//Register()
	Init()
	//SetConfig() bool
	//Activate()
	//Deactivate()
	IsActive() bool
	//Index() int
	//GetName() string
}

// GetBlock get a list of active block (by name)
func GetActiveBlocks(forced bool) []string {
	blocks := []string{}
	if !forced {
		for blockName, currentBlock := range Blocks {
			if currentBlock.IsActive() {
				blocks = append(blocks, blockName)
			}
		}
		go cache.MainCache.Set("activeBlocks", blocks, 60)
	} else {
		blocks = cache.MainCache.GetStringList("activeBlocks", 60)
	}
	return blocks
}

func init() {
	RegisteredBlocks = make(map[string]Block)
	Blocks = make(map[string]Block)
	for _, currentBlock := range Blocks {
		currentBlock.Init()
	}
	//LoadBlocks()
	//configureTemplates()
}
