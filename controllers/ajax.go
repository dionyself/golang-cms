package controllers

import (
	"encoding/json"
	"image"
	"io"
	"mime/multipart"
	"strconv"

	"github.com/dionyself/golang-cms/models"
	"github.com/dionyself/golang-cms/utils"
)

type AjaxController struct {
	BaseController
}

// generic Ajax
func (CTRL *AjaxController) PostImage() {
	form := new(models.ImageForm)
	if err := CTRL.ParseForm(form); err != nil {
		CTRL.Abort("401")
	} else {
		if form.Validate() {
			if rawFile, fileHeader, err := CTRL.GetFile("File"); err == nil && utils.Contains(utils.SuportedMimeTypes["images"], fileHeader.Header["Conten-Type"][0]) {
				defer rawFile.Close()
				newSession := utils.GetRandomString(16)
				go CTRL.uploadAndRegisterIMG(newSession, rawFile, fileHeader, form)
				response := map[string]string{"status": "started", "sessionId": newSession}
				CTRL.Data["json"] = &response
			}
		}
	}
	CTRL.ServeJSON()
}

func (CTRL *AjaxController) uploadAndRegisterIMG(sessionKey string, img io.Reader, fileHeader *multipart.FileHeader, form *models.ImageForm) {
	var croppedImg image.Image
	var targets map[string][2]int
	status := map[string]string{}
	CTRL.cache.Set(sessionKey, `{"status":"starting..."}`, 30)
	json.Unmarshal([]byte(form.Targets), &targets)
	for target, coords := range targets {
		if !utils.ContainsKey(utils.ImageSizes, target) {
			continue
		}
		croppedImg, _ = utils.CropImage(img, fileHeader.Header["Conten-Type"][0], target, coords)
		if croppedImg != nil {
			utils.UploadImage(target, croppedImg)
		}
		status[target] = "done"
		CTRL.cache.Set(sessionKey, status, 30)
	}
	user := CTRL.Data["user"].(models.User)
	newImg := new(models.Image)
	newImg.User = &user
	CTRL.db.Insert(newImg)
	CTRL.db.Insert(user)
	status["image_id"] = string(user.Id)
	CTRL.cache.Set(sessionKey, status, 30)
}

func (CTRL *AjaxController) GetImageUploadStatus() {
	imgID, err := strconv.Atoi(CTRL.Ctx.Input.Param(":id"))
	if err != nil {
		CTRL.Abort("403")
	}
	data := map[string]string{}
	if err := json.Unmarshal(CTRL.Ctx.Input.RequestBody, &data); err != nil {
		CTRL.Ctx.Output.SetStatus(400)
	}
	if status, err := CTRL.cache.GetMap(data["sessionKey"], 30); err == false {
		data = status
		data["image_id"] = string(imgID)
	}
	CTRL.Data["json"] = &data
	CTRL.ServeJSON()
}
