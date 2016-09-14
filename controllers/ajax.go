package controllers

import (
	"image"

	"github.com/dionyself/golang-cms/models"
	"github.com/dionyself/golang-cms/utils"
)

type AjaxController struct {
	BaseController
}

// generic Ajax
func (CTRL *AjaxController) PostImage() {
	form := new(models.ImageForm)
	target := ""
	var croppedImg image.Image
	if err := CTRL.ParseForm(form); err != nil {
		CTRL.Abort("401")
	} else {
		if form.Validate() {
			if rawFile, fileHeader, err := CTRL.GetFile("File"); err == nil && utils.Contains(utils.SuportedMimeTypes["images"], fileHeader.Header["Conten-Type"][0]) {
				defer rawFile.Close()
				croppedImg, _ = utils.CropImage(rawFile, fileHeader.Header["Conten-Type"][0], form.Target, [2]int{form.PivoteX, form.PivoteY})
			}
		}
	}
	if croppedImg != nil {
		response := models.UploadResultJSON{Msg: "UploadSucess"}
		CTRL.Data["json"] = &response
		user := CTRL.Data["user"].(models.User)
		url := utils.UploadImage(target, croppedImg)
		img := new(models.Image)
		img.Url = url
		img.User = &user
		CTRL.db.Insert(img)
		user.Profile.Avatar = url
		CTRL.db.Insert(user)
		CTRL.ServeJSON()
	}
}

func (CTRL *AjaxController) GetImage() {

}
