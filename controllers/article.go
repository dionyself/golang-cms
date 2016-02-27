package controllers

import (
        "github.com/dionyself/golang-cms/utils"
        "github.com/dionyself/golang-cms/models"
        "strconv"
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
	    var cats []*models.Category
        db.QueryTable("category").All(&cats)
        this.Data["Categories"] = cats
	    this.ConfigPage("article-editor.html")
	}else{
		Art := new(models.Article)
		Art.Id = ArtId
		db.Read(&Art, "Id")
		this.Data["Article"] = Art
		this.ConfigPage("article.html")
	}
}

func (this *ArticleController) Post() {
	form := utils.ArticleForm {}
	Art := new(models.Article)
	if err := this.ParseForm(&form); err != nil {
        this.Abort("401")
    }else{
    	if err := form.Validate(&Art); err != nil {
    		//  error validating form
    		this.Abort("403")
    	}
    	db := this.GetDB("default")
        if len(form.Errors) == 0 {
        	//shoudbe >0 on validation errors
            this.Data["form"] = form
            var cats []*models.Category
            db.QueryTable("category").All(&cats)
            this.Data["Categories"] = cats
	        this.ConfigPage("article-editor.html")
	    }else{
		    db.Insert(Art)
		    this.Data["Article"] = Art
		    this.ConfigPage("article.html")
	     	}
    }
}
