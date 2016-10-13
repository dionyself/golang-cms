package block

import (
	_ "github.com/astaxie/beego"
	_ "github.com/astaxie/beego/orm"
	"github.com/dionyself/golang-cms/core/lib/db"
	"github.com/dionyself/golang-cms/models"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var BlockType = "html"

type htmlBlock struct {
	Content string
	Type    string
	Name    string
	Config  []*models.BlockConfig //map[string]string
	//IsActive    bool
	IsCacheable bool
	index       int
	templates   []string
	tempate     string
}

// this maybe moved to single funct
//func (block htmlBlock) Init(siteData map[string]string, config map[string]string, content map[string]string) {
func (block htmlBlock) Init() {
	// blockModel := new(models.Block)
	// blockModel.Config = block.generateConfig(siteData map[string]string, config map[string]string)
	// blockModel.Content = block.generateContent(configModel, content map[string]string)
	// Block[config["name"]] = block.load(models.Block)
}

func (block htmlBlock) GetContent() string {
	return "test content"
}
func (block htmlBlock) GetPosition() string {
	return "1"
}

func (block htmlBlock) IsActive() bool {
	return true
}

//this couldbe reimplemented susupport mutiple themes/templates/style
// by now hardoding
func (block htmlBlock) GetTemplatePath() string {
	return "default/blocks/html_block.html"
}

func (block htmlBlock) Load(blockModel *models.Block) htmlBlock {
	//block.Type = blockModel.Type
	//block.Content = blockModel.Content
	//block.Name = blockModel.Name
	//block.Config = blockModel.Config
	return block
}

func init() {
	DB := db.MainDatabase.Orm
	htmlBlocks := []models.Block{}
	qs := DB.QueryTable("block").Filter("type", BlockType)
	qs.All(&htmlBlocks)
	toUpdate := htmlBlock{}
	RegisteredBlocks[BlockType] = toUpdate //we may to want populate with some default info
	for _, currentBlock := range htmlBlocks {
		toUpdate := new(htmlBlock)
		Blocks[currentBlock.Name] = toUpdate.Load(&currentBlock)
	}
}
