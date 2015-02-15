// @APIVersion 1.0.0
// @Title card API
// @Description 卡号管理控制器
// @Contact 3wmaocomputer@gmail.com

package controllers

import (
	"LocalServer/models"
	"fmt"
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
	card_guid := this.GetString("card")
	this.init_register(card_guid)
	this.TplNames = "card/show.tpl"
}

func (this *CardController) init_register(card string) {
}

func (this *CardController) Search() {
	search_condition := this.GetString("search_condition")

	fmt.Println(search_condition)

	familiy_model := new(models.FamiliyModel)
	//检索家庭名称
	res_fm_name := familiy_model.GetListByFamlilyName(search_condition)
	fmt.Printf("%v==\n", res_fm_name)
	this.Data["search_by_familiy"] = res_fm_name
	//检索成员姓名
	res_u_name := familiy_model.GetListByUserName(search_condition)
	fmt.Printf("%v==\n", res_fm_name)
	this.Data["search_by_user"] = res_u_name

	//检索成员手机号
	res_u_phone := familiy_model.GetListByUserPhone(search_condition)
	this.Data["search_by_phone"] = res_u_phone

	this.Data["search_condition"] = search_condition
	this.TplNames = "card/search.tpl"
}

//func init() {
//	device := models.DevicesModel{}
//	device.DemoData()
//}
