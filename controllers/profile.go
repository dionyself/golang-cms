package controllers

import (
	"strconv"

	"github.com/dionyself/golang-cms/models"
)

type ProfileController struct {
	BaseController
}

func (this *ProfileController) UserPanelView() {
	Uid := this.Ctx.Input.Param(":id")
	if this.Ctx.Input.Param(":id") == "0" {
		this.ConfigPage("user-profile.html")
	} else {
		Uid, err := strconv.Atoi(Uid)
		if err != nil {
			this.Abort("404")
		}
		this.populateProfileViewData(Uid)
		this.ConfigPage("profile-view.html")
	}
}

func (this *ProfileController) populateProfileViewData(Uid int) bool {
	db := this.GetDB()
	userView := models.User{Id: Uid}
	db.Read(&userView, "Id")
	Permissions := userView.Profile.GetPermissions(this.Data["user"].(models.User))
	// TODO : populate permissions on this.Data
	_ = Permissions
	this.Data["Profile"] = userView.Profile
	return true
}
