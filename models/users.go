package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Users struct {
	Id         uint64    `orm:"fk;auto"`
	Guid       string    `orm:"unique;size(50)"`
	Phone      string    `orm:"size(20);null"`
	Realname   string    `orm:"size(10);null"`
	Password   string    `orm:"size(50)"`
	Spell      string    `orm:"size(10);null"`
	Gender     int8      `orm:"default(1)"`
	IdCard     string    `orm:"size(20);null"`
	Picture    string    `orm:"size(100);null"`
	SchoolGuid string    `orm:"size(50)"`
	CreateTime time.Time `orm:"type(datetime) form:"ContractDate,2006-01-02 00:00:00"`
}

/**
 * 根据guid获取家长信息
 */
func (this *Users) GetUserByGuid(guid string) error {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	err := qs.Filter("guid", guid).One(this)
	return err
}

/**
 * 根据guid更新用户信息
 */
func (this *Users) UpdateUserByGuid(guid string, par map[string]interface{}) bool {
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
