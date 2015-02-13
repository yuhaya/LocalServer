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
