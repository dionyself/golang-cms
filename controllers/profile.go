package controllers

type ProfileController struct {
	BaseController
}

func (this *ProfileController) UserPanelView() {
	this.ConfigPage("user-panel.html")
}
