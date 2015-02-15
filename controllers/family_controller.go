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
	family_model.Guid = tool.Uuid()
	family_model.FirstGuid = user_guid
	family_model.SchoolGuid = school_model.GetSchoolGuid()
	family_model.Name = family_name

	familiy := models.FamiliyModel{}
	succ := familiy.AddFamily(&family_model, &user_model)

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
