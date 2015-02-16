package controllers

import (
	"LocalServer/lib"
	"LocalServer/models"
	"encoding/json"
	"fmt"
)

type GradeController struct {
	BaseController
}

func (this *GradeController) Show() {

	var list []*models.Grades
	var grade models.Grades
	count, _ := grade.Query().Count()
	if count > 0 {
		grade.Query().OrderBy("rating").All(&list)
	}
	this.Data["grade_list"] = list
	for k, v := range list {
		fmt.Println(k)
		fmt.Println(v)
	}
	this.TplNames = "grade/show.tpl"
}

func (this *GradeController) Index() {
	var list []*models.Grades
	var grade models.Grades
	count, _ := grade.Query().Count()
	if count > 0 {
		grade.Query().OrderBy("rating").All(&list)
	}
	this.Data["list"] = list
	this.TplNames = "grade/grade.tpl"
}

func (this *GradeController) Create() {
	sch := models.SchoolModel{}
	school_guid := sch.GetSchoolGuid()
	if school_guid == "" {
		this.AjaxReturnFun("1403", "学校不存在", nil)
		return
	}
	guid := lib.GetGuid()
	name := this.GetString("name")
	if name == "" {
		this.AjaxReturnFun("1404", "年级名称不能为空", nil)
		return
	}

	var g models.Grades
	g.Guid = guid
	g.Name = name
	g.Rating = setRating()
	g.SchoolGuid = school_guid
	if err := g.Insert(); err != nil {
		this.AjaxReturnFun("1405", "添加年级失败", nil)
		return
	}
	this.AjaxReturnFun("0", "Success", g)

}

func setRating() uint64 {
	var grade models.Grades
	return grade.MaxRating() + 1
}

func (this *GradeController) Delete() {
	guid := this.GetString("guid")
	var g models.Grades

	g.Guid = guid
	if err := g.Delete(guid); err != nil {
		this.AjaxReturnFun("1407", "删除年级失败", nil)
		return
	}
	this.AjaxReturnFun("0", "Success", nil)
}

func (this *GradeController) Update() {
	data := this.GetString("data")
	var g struct {
		Grades []models.Grades
	}
	json.Unmarshal([]byte(data), &g)
	var flag bool
	for _, v := range g.Grades {
		if err := v.Update(); err != nil {
			flag = true
		}
	}
	if flag {
		this.AjaxReturnFun("1408", "更新失败", nil)
		return
	}
	this.AjaxReturnFun("0", "Success", nil)
}
