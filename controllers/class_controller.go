package controllers

import "LocalServer/models"
import "LocalServer/lib"
import "github.com/astaxie/beego/orm"
import "fmt"

type ClassController struct {
	BaseController
}

func (this *ClassController) Index() {
	var class models.Classes
	guid := this.GetString("guid")
	cls, err := class.Query(guid)
	if err != nil {
		this.AjaxReturnFun("1010", "获取班级信息错误", nil)
		return
	}
	type data struct {
		Guid          string
		GradeGuid     string
		Classname     string
		TeacherGuid   string
		Teacher       string
		Assist        string
		AssistantGuid string
	}
	var datas []data
	if cls != nil {
		for i := 0; i < len(cls); i++ {
			var exist bool = false
			var d data
			d.Guid = cls[i].Guid
			d.GradeGuid = cls[i].GradeGuid
			d.Classname = cls[i].Name
			d.Teacher = cls[i].Teacher
			d.TeacherGuid = cls[i].TeacherGuid
			if len(datas) == 0 {
				datas = append(datas, d)
				i = -1
			} else {
				for j := 0; j < len(datas); j++ {
					if d.Guid == datas[j].Guid {
						exist = true
						if cls[i].Adviser == 1 {
							datas[j].Teacher = cls[i].Teacher
							datas[j].TeacherGuid = cls[i].TeacherGuid
						} else {
							datas[j].Assist = cls[i].Teacher
							datas[j].AssistantGuid = cls[i].TeacherGuid
						}
					}
				}
				if !exist {
					datas = append(datas, d)
					i = -1
				}
			}
		}
	}
	this.AjaxReturnFun("0", "获取信息成功", datas)
}

func (this *ClassController) Create() {
	guid := this.GetString("guid")
	classname := this.GetString("classname")
	teacher := this.GetString("teacher")
	assistant := this.GetString("assistant")
	sch := models.SchoolModel{}
	school_guid := sch.GetSchoolGuid()
	if school_guid == "" {
		this.AjaxReturnFun("1403", "学校不存在", nil)
		return
	}
	if guid == "" {
		this.AjaxReturnFun("1020", "请选择年级", nil)
		return
	}
	if classname == "" {
		this.AjaxReturnFun("1021", "请输入班级名称", nil)
		return
	}
	//开启事务
	failed := false
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		o.Rollback()
		this.AjaxReturnFun("1022", "保存班级信息失败", nil)
		return
	}
	//添加班级
	re_class_guid, err1 := createClass(classname, school_guid, o)
	if err1 != nil {
		o.Rollback()
		failed = true
		this.AjaxReturnFun("1022", "保存班级信息失败", nil)
		return
	}
	//绑定年级班级关系
	if err := bindGradeClass(guid, re_class_guid, school_guid, o); err != nil {
		o.Rollback()
		failed = true
		this.AjaxReturnFun("1022", "保存班级信息失败", nil)
		return
	}
	//添加教师
	if teacher != "" {
		re_teacher_guid, err2 := createTacher(teacher, school_guid, o)
		if err2 != nil {
			o.Rollback()
			failed = true
			this.AjaxReturnFun("1022", "保存班级信息失败", nil)
			return
		}
		//绑定教师班级关系
		if err := bindTeacherClass(re_teacher_guid, re_class_guid, 1, school_guid, o); err != nil {
			o.Rollback()
			failed = true
			this.AjaxReturnFun("1022", "保存班级信息失败", nil)
			return
		}
	}

	//添加助教
	if assistant != "" {
		re_assistant_guid, err3 := createTacher(assistant, school_guid, o)
		if err3 != nil {
			o.Rollback()
			failed = true
			this.AjaxReturnFun("1022", "保存班级信息失败", nil)
			return
		}
		//绑定助教班级关系
		if err := bindTeacherClass(re_assistant_guid, re_class_guid, 0, school_guid, o); err != nil {
			o.Rollback()
			failed = true
			this.AjaxReturnFun("1022", "保存班级信息失败", nil)
			return
		}
	}

	if !failed {
		o.Commit()
		this.AjaxReturnFun("0", "保存成功", nil)
		return
	}
}

//添加班级
func createClass(name string, sid string, o orm.Ormer) (class_guid string, err error) {
	guid := lib.GetGuid()
	var c models.Classes
	c.Guid = guid
	c.Name = name
	c.SchoolGuid = sid
	if err := c.InsertTrans(o); err != nil {
		return "", err
	}
	return guid, nil
}

//绑定年级班级关系
func bindGradeClass(grade_guid string, class_guid string, sid string, o orm.Ormer) error {
	var gc models.GradeClass
	gc.GradeGuid = grade_guid
	gc.ClassGuid = class_guid
	gc.SchoolGuid = sid
	if err := gc.InsertTrans(o); err != nil {
		return err
	}
	return nil
}

//添加教师
func createTacher(name string, sid string, o orm.Ormer) (teacher_guid string, err error) {
	guid := lib.GetGuid()
	var t models.Teachers
	t.Guid = guid
	t.Name = name
	t.Phone = ""
	t.Gender = 0
	t.SchoolGuid = sid
	if err := t.InsertTrans(o); err != nil {
		return "", err
	}
	return guid, nil
}

//绑定教师班级关系
func bindTeacherClass(teacher_guid string, class_guid string, adviser int8, sid string, o orm.Ormer) error {
	var tc models.ClassTeacher
	tc.ClassGuid = class_guid
	tc.TeacherGuid = teacher_guid
	tc.SchoolGuid = sid
	tc.Adviser = adviser
	if err := tc.InsertTrans(o); err != nil {
		return err
	}
	return nil
}

func (this *ClassController) Delete() {
	grade_guid := this.GetString("grade_guid")
	class_guid := this.GetString("class_guid")
	teacher_guid := this.GetString("teacher_guid")
	assistant_guid := this.GetString("assistant_guid")
	fmt.Println("teacher_guid" + teacher_guid)
	fmt.Println("assistant_guid" + assistant_guid)
	if grade_guid == "" || class_guid == "" {
		this.AjaxReturnFun("1031", "请求参数错误", nil)
		return
	}
	//开启事务
	failed := false
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		o.Rollback()
		this.AjaxReturnFun("1030", "删除失败", nil)
		return
	}
	//删除年级班级关系
	if err := deleteGradeClass(grade_guid, class_guid, o); err != nil {
		o.Rollback()
		failed = true
		this.AjaxReturnFun("1030", "删除失败", nil)
		return
	}
	//删除班级信息
	if err := deleteClass(class_guid, o); err != nil {
		o.Rollback()
		failed = true
		this.AjaxReturnFun("1030", "删除失败", nil)
		return
	}
	//删除教师班级关系
	if err := deleteTeacherClass(class_guid, o); err != nil {
		o.Rollback()
		failed = true
		this.AjaxReturnFun("1030", "删除失败", nil)
		return
	}
	//删除班主任
	if teacher_guid != "" {
		if err := deleteTeacher(teacher_guid, o); err != nil {
			o.Rollback()
			failed = true
			this.AjaxReturnFun("1030", "删除失败", nil)
			return
		}
	}

	//删除助教
	if assistant_guid != "" {
		if err := deleteTeacher(assistant_guid, o); err != nil {
			o.Rollback()
			failed = true
			this.AjaxReturnFun("1030", "删除失败", nil)
			return
		}
	}

	if !failed {
		o.Commit()
		this.AjaxReturnFun("0", "保存成功", nil)
		return
	}
}

//删除年级班级关系
func deleteGradeClass(grade_guid string, class_guid string, o orm.Ormer) error {
	var gc models.GradeClass
	if err := gc.DeleteTrans(grade_guid, class_guid, o); err != nil {
		return err
	}
	return nil
}

//删除班级信息
func deleteClass(class_guid string, o orm.Ormer) error {
	var cl models.Classes
	if err := cl.DeleteTrans(class_guid, o); err != nil {
		return err
	}
	return nil
}

//删除教师班级关系
func deleteTeacherClass(class_guid string, o orm.Ormer) error {
	var tc models.ClassTeacher
	if err := tc.DeleteTrans(class_guid, o); err != nil {
		return err
	}
	return nil
}

//删除教师信息
func deleteTeacher(guid string, o orm.Ormer) error {
	var t models.Teachers
	if err := t.DeleteTrans(guid, o); err != nil {
		return err
	}
	return nil
}
