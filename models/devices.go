package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Devices struct {
	Id          uint64 `orm:"fk;auto"`
	Guid        string `orm:"unique;size(50)"`
	Device      string `orm:"size(50)"`
	Kind        int8   `orm:"default(0)"`
	Vmp         string `orm:"size(10)"`
	SchoolGuid  string `orm:"size(50)"`
	Group       int8   `orm:"default(0)"`
	Description string `orm:"size(255)"`
	Status      int8   `orm:"default(1)"`
	Enabled     int8   `orm:"default(1)"`
}

type DevicesModel struct {
}

func (this *DevicesModel) List() ([]*Devices, int64, error) {
	o := orm.NewOrm()
	device_model := new(Devices)
	qs := o.QueryTable(device_model)
	var devices_list []*Devices
	num, err := qs.Filter("Enabled", 1).All(&devices_list)
	return devices_list, num, err
}

//测试更新
func (this *DevicesModel) UpdateAll() {
	o := orm.NewOrm()
	device_model := new(Devices)
	num, err := o.QueryTable(device_model).Filter("Enabled", 1).Update(orm.Params{
		"Vmp": "armand",
	})
	fmt.Printf("Affected Num: %s, %s", num, err)
}
