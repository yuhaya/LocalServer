package controllers

import (
    "LocalServer/models"
    "LocalServer/lib"
    "strconv"
    "time"

 )

type SchoolController struct {
    BaseController
}

func (this *SchoolController) Index() {
    //查询是否存在学校
    var school models.Schools
    var sm models.SchoolModel
    guid := sm.GetSchoolGuid()
    if guid!="" {
        //根据guid获取学校信息
        school.Query(guid).One(&school)
    }
    this.Data["school"]=school
    this.TplNames="school/index.tpl"
}

func (this *SchoolController) Create() {
    guid := this.GetString("guid")
    name := this.GetString("name")
    province := this.GetString("province")
    city := this.GetString("city")
    county := this.GetString("county")
    address := this.GetString("address")
    door := this.GetString("door")
    var i, _ = strconv.Atoi(door)
    var school models.Schools
    school.Guid = guid
    school.Name=name
    school.Province = province
    school.City = city
    school.County = county
    school.Address = address
    school.DoorNum = int8(i)
    school.UpdateTime = time.Now()
    school.Enabled = 1
    if name!="" {
        if guid!="" {
            if err:=school.Update();err!=nil{
                this.AjaxReturnFun("1404", "保存失败", nil)
                return
            }
        }else {
            school.Guid = lib.GetGuid()
            if err := school.Insert(); err!=nil {
                this.AjaxReturnFun("1403", "保存失败", nil)
                return
            }
        }
        this.AjaxReturnFun("0", "保存成功", nil)
    }

}