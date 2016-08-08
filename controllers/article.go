package controllers

import (
	"fmt"
	"strconv"

	"github.com/dionyself/golang-cms/models"
)

// ArticleController ...
type ArticleController struct {
	BaseController
}

// Get Displays Article by id
func (CTRL *ArticleController) Get() {
	ArtID, err := strconv.Atoi(CTRL.Ctx.Input.Param(":id"))
	if err != nil {
		CTRL.Abort("403")
	}
	db := CTRL.GetDB("default")
	if ArtID == 0 {
		CTRL.Data["form"] = models.ArticleForm{}
		cats := new([]*models.Category)
		db.QueryTable("category").All(cats)
		CTRL.Data["Categories"] = *cats
		CTRL.ConfigPage("article-editor.html")
	} else {
		Art := new(models.Article)
		Art.Id = ArtID
		db.Read(Art, "Id")
		CTRL.Data["Article"] = Art
		CTRL.ConfigPage("article.html")
	}
}

// Post create/update article
func (CTRL *ArticleController) Post() {
	form := new(models.ArticleForm)
	Art := new(models.Article)
	if err := CTRL.ParseForm(form); err != nil {
		CTRL.Abort("401")
	} else {
		db := CTRL.GetDB()
		if !form.Validate() {
			CTRL.Data["form"] = form
			cats := new([]*models.Category)
			db.QueryTable("category").All(cats)
			CTRL.Data["Categories"] = cats
			CTRL.ConfigPage("article-editor.html")
			for key, msg := range form.InvalidFields {
				fmt.Println(key, msg)
			}
		} else {
			cat := new(models.Category)
			cat.Id = form.Category
			db.Read(cat, "Id")
			Art.Category = cat
			user := CTRL.Data["user"].(models.User)
			Art.User = &user
			Art.Title = form.Title
			Art.Content = form.Content
			Art.AllowComments = form.AllowComments
			db.Insert(Art)
			CTRL.Data["Article"] = Art
			CTRL.ConfigPage("article.html")
		}
	}
}
