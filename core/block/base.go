package block

import (
	_ "github.com/dionyself/beego"
	_ "github.com/dionyself/beego/orm"
	"github.com/dionyself/golang-cms/core/lib/cache"
	"github.com/dionyself/golang-cms/core/lib/db"
	"github.com/dionyself/golang-cms/models"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type Block interface {
	GetTemplatePath() string
	GetContent() map[string]string
	GetPosition() int
	GetBlockType() string
	//IsContentCacheable() bool
	//GetConfig() models.BlockConfig
	//Register()
	Init()
	Load(*models.Block) Block
	//SetConfig() bool
	//Activate()
	//Deactivate()
	IsActive() bool
	//Index() int
	//GetName() string
}

// Blocks map["block_folder"]["style1", "style2" ...]
var Blocks map[string]Block
var RegisteredBlocks map[string]Block

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

// initBlock use this to initialize your block
func initBlock(blockToInit Block) {
	BlockType := blockToInit.GetBlockType()
	DB := db.MainDatabase.Orm
	byTypeBlocks := []models.Block{}
	qs := DB.QueryTable("block").Filter("type", BlockType)
	qs.All(&byTypeBlocks)
	RegisteredBlocks[BlockType] = blockToInit //we may to want populate with some default info
	for _, currentBlock := range byTypeBlocks {
		Blocks[currentBlock.Name] = blockToInit.Load(&currentBlock)
	}
}

func init() {
	RegisteredBlocks = make(map[string]Block)
	Blocks = make(map[string]Block)
	for _, currentBlock := range Blocks {
		currentBlock.Init()
	}
}
