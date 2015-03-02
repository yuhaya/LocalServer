// @APIVersion 1.0.0
// @Title demo API
// @Description ${PROJECT_NAME} | this is demo description
// @Contact mao | 3wmaocomputer@gmail.com
// @Date 2/28/15

package controllers

import (
	"LocalServer/models"
	"fmt"
	"time"
)

type PointController struct {
	BaseController
}

func (this *PointController) Create() {

	card := this.GetString("card")
	time_val := this.GetString("time")
	type_val, err_type := this.GetInt8("type")
	auto_val, err_auto := this.GetInt8("auto")

	device := this.GetString("device")
	kind, _ := this.GetInt8("kind")
	vmp := this.GetString("vmp")

	mode := this.BaseController.Mode()
	if mode == 1 {
		if auto_val == 1 {
			this.AjaxReturnFun("0", "当前处于注册模式中，无法录入到校信息!", nil)
			return
		} else {
			println(err_type, err_auto, card, time_val)
			urlmsg := make(map[string]string)
			urlmsg["返回上一页"] = "javascript:history.go(-1)"
			this.OutputMsg("当前处于注册模式中，无法录入到校信息!", urlmsg)
			return
		}
	}

	if err_type != nil || err_auto != nil || card == "" || time_val == "" {
		if auto_val == 1 {
			this.AjaxReturnFun("0", "参数信息错误!", nil)
			return
		} else {
			println(err_type, err_auto, card, time_val)
			urlmsg := make(map[string]string)
			urlmsg["返回上一页"] = "javascript:history.go(-1)"
			this.OutputMsg("参数信息错误", urlmsg)
			return
		}
	}

	flag := true
	msg := ""
	if device != "" {
		fmt.Println("sdsssssssssssssssssssssssssssss%s", device)
		var de models.Devices
		de.Device = device
		de.Kind = kind
		flag_tmp, msg_tmp := de.UpdateVmp(vmp)
		flag = flag_tmp
		msg = msg_tmp
	}

	var att models.Attendances
	att.Auto = auto_val
	att.Card = card
	val_time, err_time := time.Parse("2006-01-02 15:04:05", time_val)
	if err_time != nil {
		if auto_val == 1 {
			this.AjaxReturnFun("0", "时间解析错误!", nil)
			return
		} else {
			urlmsg := make(map[string]string)
			urlmsg["返回上一页"] = "javascript:history.go(-1)"
			this.OutputMsg("时间解析错误", urlmsg)
			return
		}
	}
	att.Time = val_time
	school_model := models.SchoolModel{}
	att.SchoolGuid = school_model.GetSchoolGuid()
	att.Type = type_val
	flag2, _, msg2 := att.Insert()

	if auto_val == 1 {
		//自动输入
		if !(flag && flag2) {
			this.AjaxReturnFun("0", msg+"/"+msg2, nil)
		}
		var mc models.MemberCard
		mc.Card = card
		res := mc.GetStuMsg()
		res.Mode = mode
		if res != nil {
			this.AjaxReturnFun("1", "success", res)
			return
		} else {
			this.AjaxReturnFun("0", "未检查到卡号对应的学生信息!", nil)
			return
		}

	} else {
		//手动输入
		urlmsg := make(map[string]string)
		urlmsg["返回上一页"] = "javascript:history.go(-1)"
		if flag && flag2 {
			this.OutputMsg("添加成功！", urlmsg)
		} else {
			this.OutputMsg("添加失败！"+msg+"/"+msg2, urlmsg)
		}

	}

}

func (this *PointController) CreateShow() {
	this.TplNames = "point/createshow.tpl"
}
