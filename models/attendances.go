package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Attendances struct {
	Id         uint64    `orm:"fk;auto"`
	Card       string    `orm:"size(50)"`
	Time       time.Time `orm:"type(datetime)"`
	Type       int8      `orm:"default(0)"`
	SchoolGuid string    `orm:"size(50)"`
	Auto       int8      `orm:"default(1)"`
}

func (this *Attendances) Insert() (bool, int64, string) {
	o := orm.NewOrm()
	id, err := o.Insert(this)
	if err == nil {
		return true, id, ""
	} else {
		return false, 0, err.Error()
	}
}
