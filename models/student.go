package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Students struct {
	Id          uint64    `orm:"fk;auto"`
	Guid        string    `orm:"unique;size(50)"`
	Sid         string    `orm:"size(50);null"`
	Realname    string    `orm:"size(10)"`
	Spell       string    `orm:"size(10)"`
	Gender      int8      `orm:"default(1)"`
	Picture     string    `orm:"size(255);null"`
	Birthday    time.Time `orm:"default(0000-00-00);type(date);" form:"Birthday,2006-01-02"`
	School_guid string    `orm:"size(50);null"`
	Grade_guid  string    `orm:"size(50);null"`
	Class_guid  string    `orm:"size(20);null"`
	Enrol_time  time.Time `orm:"default(0000-00-00);type(date);" form:"Enrol_time,2006-01-02"`
	Create_time time.Time `orm:"type(datetime)"`
}

func (this *Students) GetStudentByGuid(guid string) error {
	err := orm.NewOrm().QueryTable(this).Filter("guid", guid).One(this)
	return err
}

/**
 * 根据guid更新学生信息
 */
func (this *Students) UpdateStuByGuid(guid string, par map[string]interface{}) bool {
	o := orm.NewOrm()
	num, err := o.QueryTable(this).Filter("guid", guid).Update(par)
	if err != nil {
		//日志记录
	}
	if num > 0 && err == nil {
		return true
	} else {
		return false
	}
}

/**
 * 删除学生
 */
func (this *Students) DeleteStu(guid string, fm string) bool {
	o := orm.NewOrm()
	err := o.Begin()

	_, err1 := o.QueryTable(this).Filter("guid", guid).Delete()
	_, err2 := o.QueryTable(new(FamilyMember)).Filter("member_guid", guid).Filter("family_guid", fm).Delete()
	_, err3 := o.QueryTable(new(FamilyRelation)).Filter("user_guid", guid).Delete()
	_, err4 := o.QueryTable(new(CardReceiver)).Filter("guid", guid).Delete()
	_, err5 := o.QueryTable(new(MemberCard)).Filter("guid", guid).Delete()

	var err6 error
	var flag bool
	if err == nil && err1 == nil && err2 == nil && err3 == nil && err4 == nil && err5 == nil {
		flag = true
		err6 = o.Commit()
	} else {
		flag = false
		err6 = o.Rollback()
	}
	if err == nil && err6 == nil && flag {
		return true
	} else {
		return false
	}
}
