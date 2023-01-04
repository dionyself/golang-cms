package block

import (
	"encoding/json"

	_ "github.com/beego/beego/v2/client/orm"
	_ "github.com/beego/beego/v2/server/web"
	"github.com/dionyself/golang-cms/models"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type htmlBlock struct {
	content     map[string]string
	Type        string
	Name        string
	Config      []*models.BlockConfig //map[string]string
	isActive    bool
	IsCacheable bool
	index       int
	templates   []string
	tempate     string
	position    int
}

// this maybe moved to single funct
// func (block htmlBlock) Init(siteData map[string]string, config map[string]string, content map[string]string) {
func (block htmlBlock) Init() {
	// blockModel := new(models.Block)
	// blockModel.Config = block.generateConfig(siteData map[string]string, config map[string]string)
	// blockModel.Content = block.generateContent(configModel, content map[string]string)
	// Block[config["name"]] = block.load(models.Block)
}

// GetBlockType yup this is hardcoded an this the way to do it
func (block htmlBlock) GetBlockType() string {
	return "html"
}

func (block htmlBlock) GetContent() map[string]string {
	return block.content
}
func (block htmlBlock) GetPosition() int {
	return block.position
}

func (block htmlBlock) IsActive() bool {
	return block.isActive
}

// this couldbe reimplemented support mutiple themes/templates/style
// by now hardoding
func (block htmlBlock) GetTemplatePath() string {
	return "default/blocks/html_block.html"
}

func (block htmlBlock) Load(blockModel *models.Block) Block {
	//block.Type = blockModel.Type
	json.Unmarshal([]byte(blockModel.Content), &block.content)
	block.Name = blockModel.Name
	//block.Config = blockModel.Config
	block.position = blockModel.Position
	block.isActive = blockModel.IsActive
	return block
}

func init() {
	initBlock(htmlBlock{})
}
