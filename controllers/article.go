package controllers

import (

)

type ArticleController struct {
	BaseController
}

func (this *ArticleController) Get() {
	this.ConfigPage("article-editor.html")
}

