package controllers

import (
	"LocalServer/models"
)

type CardController struct {
	BaseController
}

func (this *CardController) Index() {
	this.TplNames = "card/index.tpl"
}

func (this *CardController) Manager() {
	device_model := models.DevicesModel{}
	device_list, num, err := device_model.List()
	device_model.UpdateAll()
	this.Data["device_list"] = device_list
	this.Data["num"] = num
	if err == nil {

		this.Data["err"] = ""
	} else {

		this.Data["err"] = err.Error()
	}
	this.TplNames = "card/manager.tpl"
}

func (this *CardController) Show() {
	//	card_guid = this.GetString("card")
}
