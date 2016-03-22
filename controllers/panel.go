package controllers

type UserPanelController struct {
	BaseController
}

func (CTRL *UserPanelController) MainView() {
	CTRL.ConfigPage("user-panel.html")
}

type VendorPanelController struct {
	BaseController
}
type AdminPanelController struct {
	BaseController
}
