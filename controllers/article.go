package controllers

import (
        "github.com/dionyself/golang-cms/utils"
        "github.com/dionyself/golang-cms/models"
)

type ArticleController struct {
	BaseController
}

func (this *ArticleController) Get() {
	this.ConfigPage("article-editor.html")
}

func (this *ArticleController) Post() {
	form := utils.ArticleForm {}
	if err := this.ParseForm(&form); err != nil {
        this.Abort("401")
    }else{
    	if err := form.Validate(); err != nil {
    		//  error validating form
    		this.Abort("403")
    	}
        if len(form.Errors) == 0 {
        	//shoudbe >0 on validation errors
            this.Data["form"] = form
            cat := models.Category{}
            this.Data["Categories"] = cat
	        this.ConfigPage("article-editor.html")
	    }else{
	     	this.ConfigPage("article.html")
	     	}
    }
}
