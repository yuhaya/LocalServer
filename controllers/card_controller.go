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
	mode := this.BaseController.Mode()
	if mode == 1 {
		this.Redirect("/card/list", 302)
	} else {
		this.TplNames = "card/index.tpl"
	}
}

func (this *CardController) Manager() {
	this.CacheObj.Put("RUNMODE", 1, 90000000000)
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
	res_fm_name, num1 := familiy_model.GetListByFamlilyName(search_condition)
	fmt.Printf("%v==\n", res_fm_name)
	this.Data["search_by_familiy"] = res_fm_name
	this.Data["search_by_familiy_num"] = num1
	//检索成员姓名
	res_u_name, num2 := familiy_model.GetListByUserName(search_condition)
	fmt.Printf("%v==\n", res_fm_name)
	this.Data["search_by_user"] = res_u_name
	this.Data["search_by_user_num"] = num2

	//检索成员手机号
	res_u_phone, num3 := familiy_model.GetListByUserPhone(search_condition)
	fmt.Printf("%#v==\n", res_u_phone)
	this.Data["search_by_phone"] = res_u_phone
	this.Data["search_by_phone_num"] = num3

	this.Data["search_condition"] = search_condition
	this.TplNames = "card/search.tpl"
}

//func init() {
//	device := models.DevicesModel{}
//	device.DemoData()
//}
/**
 * 显示家庭成员
 */
func (this *CardController) MemeberList() {
	family_guid := this.GetString("family_guid")
	fm := models.FamiliyModel{}
	members, main_guid := fm.GetMembersByGuid(family_guid)

	guid_slice := make([]string, len(members["users"])+len(members["stus"]))

	count := 0
	for _, val := range members["users"] {
		guid_slice[count] = val.Guid
		count++
	}

	for _, val := range members["stus"] {
		guid_slice[count] = val.Guid
		count++
	}
	var mc models.MemberCard
	this.Data["cards"] = mc.GetAllCardsByGuids(guid_slice)
	this.Data["members"] = members
	this.Data["main_guid"] = main_guid
	this.Data["family_guid"] = family_guid

	this.TplNames = "card/member.tpl"

}

/**
 * 添加卡号
 */
func (this *CardController) Add() {

	guid := this.GetString("guid")
	card := this.GetString("card")
	//	family_guid := this.GetString("family_guid")
	type_num, _ := this.GetInt8("type")
	urlmsg := make(map[string]string)
	urlmsg["返回上一页"] = this.Referer()

	if guid != "" && card != "" {
		var mem_card models.MemberCard
		mem_card.Guid = guid
		mem_card.Card = card
		mem_card.Type = type_num
		school_model := models.SchoolModel{}
		mem_card.SchoolGuid = school_model.GetSchoolGuid()

		flag, _, _ := mem_card.InsertOne()

		if flag {
			this.OutputMsg("添加成功! ", urlmsg)
		} else {
			this.OutputMsg("添加失败,联系管理员! ", urlmsg)
		}
	} else {
		this.OutputMsg("数据提交有误! ", urlmsg)
	}
}

/**
 * 删除卡号
 */
func (this *CardController) Del() {

	guid := this.GetString("guid")
	card := this.GetString("card")
	//	fm := this.GetString("family_guid")
	var mem_card models.MemberCard
	mem_card.Card = card
	mem_card.Guid = guid
	flag := mem_card.Delete()

	urlmsg := make(map[string]string)
	urlmsg["返回上一页"] = this.Referer()

	if flag {
		this.OutputMsg("删除成功! ", urlmsg)
	} else {
		this.OutputMsg("删除失败! ", urlmsg)
	}

}
