package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type CardReceiver struct {
	Id         uint64 `orm:"fk;auto"`
	Card       string `orm:"size(50)"`
	Guid       string `orm:"size(50)"`
	Type       int8
	SchoolGuid string `orm:"size(50)"`
}

func (this *CardReceiver) Add() (bool, string) {
	o := orm.NewOrm()
	err := o.Read(this, "card", "guid")
	if err == nil {
		return false, "该卡的通知对象已经有该家庭成员!"
	} else {
		_, err2 := o.Insert(this)
		if err2 == nil {
			return true, "success!"
		} else {
			return false, err2.Error()
		}
	}
}

func (this *CardReceiver) Del() (bool, string) {
	o := orm.NewOrm()
	_, err := o.QueryTable(this).Filter("card", this.Card).Filter("guid", this.Guid).Delete()
	fmt.Printf("%#v", this)
	if err == nil {
		return true, "success!"
	} else {
		return false, err.Error()
	}
}
