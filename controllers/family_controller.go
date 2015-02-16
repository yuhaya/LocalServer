// @APIVersion 1.0.0
// @Title demo API
// @Description ${PROJECT_NAME} | this is demo description
// @Contact mao | 3wmaocomputer@gmail.com
// @Date 2/14/15

package controllers

import (
	"LocalServer/models"
	"LocalServer/tool"
	"fmt"
	"github.com/astaxie/beego/utils/pagination"
	"time"
)

type FamilyController struct {
	BaseController
}

func (this *FamilyController) Index() {
	page, err := this.GetInt("page")
	if err != nil || page < 1 {
		page = 1
	}
	familiy_model := models.FamiliyModel{}

	postsPerPage := PAGE_NUM
	fmt.Printf("%d====\n", familiy_model.CountAll())
	paginator := pagination.SetPaginator(this.Ctx, postsPerPage, familiy_model.CountAll())
	data, num := familiy_model.GetAllPage(paginator.Offset(), PAGE_NUM)
	this.Data["paginator"] = paginator
	this.Data["list"] = data
	this.Data["num"] = num
	this.TplNames = "family/index.tpl"
}

func (this *FamilyController) Add() {

	this.TplNames = "family/add.tpl"
}

func (this *FamilyController) AddSubmit() {

	family_name := this.GetString("family_name")
	name := this.GetString("name")
	phone := this.GetString("phone")
	genner, _ := this.GetInt8("gender")
	id_card := this.GetString("id_card")
	password := this.GetString("password")
	user_model := models.Users{}
	user_model.Gender = genner
	user_model.Realname = name
	user_guid := tool.Uuid()
	user_model.Guid = user_guid

	school_model := models.SchoolModel{}
	user_model.SchoolGuid = school_model.GetSchoolGuid()
	user_model.CreateTime = time.Now()
	user_model.IdCard = id_card
	user_model.Password = password
	user_model.Phone = phone

	family_model := models.Families{}
	fm_guid := tool.Uuid()
	family_model.Guid = fm_guid
	family_model.FirstGuid = user_guid
	family_model.SchoolGuid = school_model.GetSchoolGuid()
	family_model.Name = family_name

	family_member := models.FamilyMember{}
	family_member.SchoolGuid = school_model.GetSchoolGuid()
	family_member.FamilyGuid = fm_guid
	family_member.MemberGuid = user_guid
	family_member.Type = 1

	familiy := models.FamiliyModel{}
	succ := familiy.AddFamily(&family_model, &user_model, &family_member)

	urlmsg := make(map[string]string)
	urlmsg["返回上一页"] = "javascript:history.go(-1)"
	urlmsg["回到列表页"] = "/family"
	urlmsg["重新添加一个"] = "/family/add"
	if succ {
		this.OutputMsg("添加成功！", urlmsg)
	} else {
		this.OutputMsg("添加失败！", urlmsg)
	}

}

/**
 * 编辑家庭
 */
func (this *FamilyController) EditFamily() {

	family_guid := this.GetString("guid")
	family_model := models.FamiliyModel{}
	fm := family_model.GetFamilyByGuid(family_guid)
	this.Data["fm"] = fm
	this.TplNames = "family/edit.tpl"
}

/**
 * 编辑家庭提交
 */
func (this *FamilyController) EditSubmit() {

	family_guid := this.GetString("guid")
	family_name := this.GetString("family_name")
	fm := models.FamiliyModel{}
	flag := fm.UpdateFamilyNameByGuid(family_guid, family_name)

	urlmsg := make(map[string]string)
	urlmsg["返回上一页"] = "javascript:history.go(-1)"
	urlmsg["回到列表页"] = "/family"

	if flag {
		this.OutputMsg("修改成功！", urlmsg)
	} else {
		this.OutputMsg("修改失败！", urlmsg)
	}
}

/**
 * 管理家庭成员
 */
func (this *FamilyController) ManagerMember() {
	family_guid := this.GetString("guid")
	fm := models.FamiliyModel{}
	members, main_guid := fm.GetMembersByGuid(family_guid)
	this.Data["members"] = members
	this.Data["main_guid"] = main_guid
	this.TplNames = "family/member.tpl"
}
