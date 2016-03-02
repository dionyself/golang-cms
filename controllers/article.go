package controllers

import (
	"fmt"
	"strconv"

	"github.com/dionyself/golang-cms/models"
)

type ArticleController struct {
	BaseController
}

func (this *ArticleController) Get() {
	ArtId, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err != nil {
		this.Abort("403")
	}
	db := this.GetDB("default")
	if ArtId == 0 {
		this.Data["form"] = models.ArticleForm{}
		var cats []*models.Category
		db.QueryTable("category").All(&cats)
		this.Data["Categories"] = cats
		this.ConfigPage("article-editor.html")
	} else {
		Art := new(models.Article)
		Art.Id = ArtId
		db.Read(&Art, "Id")
		this.Data["Article"] = Art
		this.ConfigPage("article.html")
	}
}

func (this *ArticleController) Post() {
	form := models.ArticleForm{}
	Art := new(models.Article)
	if err := this.ParseForm(&form); err != nil {
		this.Abort("401")
	} else {
		db := this.GetDB()
		if !form.Validate() {
			this.Data["form"] = form
			var cats []*models.Category
			db.QueryTable("category").All(&cats)
			this.Data["Categories"] = cats
			this.ConfigPage("article-editor.html")
			for key, msg := range form.InvalidFields {
				fmt.Println(key, msg)
			}
		} else {
			cat := new(models.Category)
			cat.Id = form.Category
			db.Read(cat, "Id")
			Art.Category = cat
			user := this.Data["user"].(models.User)
			Art.User = &user
			Art.Title = form.Title
			Art.Content = form.Content
			Art.AllowComments = form.AllowComments
			db.Insert(Art)
			this.Data["Article"] = Art
			this.ConfigPage("article.html")
		}
	}
}
