package controllers

// UserPanelController ...
type UserPanelController struct {
	BaseController
}

// MainView ...
func (CTRL *UserPanelController) MainView() {
	CTRL.ConfigPage("user-panel.html")
}

// VendorPanelController ...
type VendorPanelController struct {
	BaseController
}

// AdminPanelController ...
type AdminPanelController struct {
	BaseController
}
