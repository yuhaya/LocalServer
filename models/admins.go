package models

import (
	"github.com/astaxie/beego/validation"
	"strings"
	"time"
)

type Admins struct {
	Id         uint64    `orm:"fk;auto"`
	Guid       string    `orm:"unique;size(50)" valid:"Required"`
	Username   string    `orm:"size(50)" valid:"Required"`
	Password   string    `orm:"size(50)" valid:"Required"`
	Realname   string    `orm:"size(50);null"`
	SchoolGuid string    `orm:"size(50)" valid:"Required"`
	CreateTime time.Time `orm:"type(datetime)" valid:"Required"`
	Super      byte      `orm:"default(0)" valid:"Required;Range(0, 1)"`
	Enabled    byte      `orm:"default(1)" valid:"Required;Range(0, 1)""`
}

func (this *Admins) Valid(v *validation.Validation) {
	if strings.Index(this.Username, "admin") != -1 {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		v.SetError("Name", "名称里不能含有 admin")
	}
}
