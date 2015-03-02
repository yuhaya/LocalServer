package controllers

import (
	_ "LocalServer/models"
)

type MainController struct {
	BaseController
}

func (this *MainController) Main() {
	this.TplNames = "main/frame.tpl"
}

func (this *MainController) TopBody() {

	this.Data["Name"] = "管理员"
	this.TplNames = "main/topbody.tpl"
}

func (this *MainController) LeftBody() {
	this.TplNames = "main/leftbody.tpl"
}

func (this *MainController) RightBody() {
	this.Layout = "main/layout.tpl"
	this.TplNames = "main/rightbody.tpl"
}
func (this *MainController) Mode() {
	mode := this.BaseController.Mode()
	this.AjaxReturnFun("1", "success", mode)
}

func (this *MainController) Card() {
	card := this.BaseController.RegCard()
	this.AjaxReturnFun("1", "success", card)
}

func (this *MainController) SetMode() {
	err := this.BaseController.SetMode()

	if err == nil {
		this.AjaxReturnFun("1", "success", nil)
	} else {
		this.AjaxReturnFun("0", err.Error(), nil)
	}
}
