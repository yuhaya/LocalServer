// @APIVersion 1.0.0
// @Title demo API
// @Description ${PROJECT_NAME} | this is demo description
// @Contact mao | 3wmaocomputer@gmail.com
// @Date 2/14/15

package controllers

import (
	"LocalServer/models"
	"LocalServer/tool"
	"encoding/json"
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
		this.OutputMsg("修改失败! ", urlmsg)
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
	this.Data["users"] = members["users"]
	this.Data["main_guid"] = main_guid
	this.Data["family_guid"] = family_guid
	this.TplNames = "family/member.tpl"
}

/**
 * 添加家庭成员显示页面
 */
func (this *FamilyController) AddMemberShow() {
	memeber_type := this.GetString("memeber_type")
	family_guid := this.GetString("family_guid")
	this.Data["memeber_type"] = memeber_type
	this.Data["family_guid"] = family_guid

	if memeber_type == "stu" {
		//添加学生
		grade_class, grade := models.GetAllGradeClass()
		grade_class_json, err := json.Marshal(grade_class)
		if err == nil && len(grade_class) != 0 {
			this.Data["grade_class_json"] = string(grade_class_json)
		} else {
			this.Data["grade_class_json"] = "{}"
		}
		fmt.Printf("\n%v======%s\n", grade_class, grade_class_json)

		this.Data["grades"] = grade
		this.TplNames = "family/addstu.tpl"
	} else {
		//添加家长
		this.TplNames = "family/adduser.tpl"
	}
}

/**
 * 添加家庭成员
 */
func (this *FamilyController) AddMember() {

	memeber_type := this.GetString("memeber_type")
	family_guid := this.GetString("family_guid")

	urlmsg := make(map[string]string)
	urlmsg["返回上一页"] = "javascript:history.go(-1)"

	if memeber_type == "stu" {

		urlmsg["再添加一个"] = "/family/members/user/add?memeber_type=stu&family_guid=" + family_guid
		//添加学生

		stu_model := models.Students{}
		if err := this.ParseForm(&stu_model); err != nil {
			//handle error

			this.OutputMsg(err.Error(), urlmsg)
		} else {
			stu_model.Create_time = time.Now()
			stu_guid := tool.Uuid()
			stu_model.Guid = stu_guid
			school_model := models.SchoolModel{}
			stu_model.School_guid = school_model.GetSchoolGuid()

			family_member := models.FamilyMember{}
			family_member.FamilyGuid = family_guid
			family_member.MemberGuid = stu_guid
			family_member.Type = 0

			family_model := models.FamiliyModel{}
			flag := family_model.AddStu(&stu_model, &family_member)

			if flag {
				this.OutputMsg("添加成功", urlmsg)
			} else {
				this.OutputMsg("添加失败", urlmsg)
			}
		}

	} else {

		urlmsg["再添加一个"] = "/family/members/user/add?memeber_type=user&family_guid=" + family_guid
		//添加家长
		user_model := models.Users{}

		if err := this.ParseForm(&user_model); err != nil {
			//handle error

			this.OutputMsg(err.Error(), urlmsg)
		} else {
			user_model.CreateTime = time.Now()
			user_guid := tool.Uuid()
			user_model.Guid = user_guid
			school_model := models.SchoolModel{}
			user_model.SchoolGuid = school_model.GetSchoolGuid()

			family_member := models.FamilyMember{}
			family_member.FamilyGuid = family_guid
			family_member.MemberGuid = user_guid
			family_member.Type = 1

			family_model := models.FamiliyModel{}
			flag := family_model.AddUser(&user_model, &family_member)

			if flag {
				this.OutputMsg("添加成功", urlmsg)
			} else {
				this.OutputMsg("添加失败", urlmsg)
			}
		}
	}
}

/**
 * 编辑家长
 */
func (this *FamilyController) EditUser() {

}

/**
 * 编辑家长
 */
func (this *FamilyController) ShowUser() {

}

/**
 * 编辑学生
 */
func (this *FamilyController) EditStu() {

}

/**
 * 查看学生详情
 */
func (this *FamilyController) ShowStu() {

}

/**
 * set main parent
 */
func (this *FamilyController) SetMainUser() {

}
