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
