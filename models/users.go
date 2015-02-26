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

/**
 * 删除用户
 */
func (this *Users) DeleteUser(guid string, fm string) bool {
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
