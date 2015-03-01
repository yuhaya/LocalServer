package models

import (
	"LocalServer/tool"
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

func (this *Devices) UpdateVmp(vmp string) (bool, string) {
	o := orm.NewOrm()
	err := o.Read(this)
	if err == nil {
		this.Vmp = vmp
		if _, err := o.Update(this); err == nil {
			return true, ""
		} else {
			return false, "更新失败"
		}
	} else {
		return false, "未找到相关记录"
	}
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

//生成测试数据
func (this *DevicesModel) DemoData() {
	o := orm.NewOrm()
	device_model := new(Devices)
	device_model.Guid = tool.Uuid()
	device_model.Device = "127.150.10.1"
	device_model.Kind = 1
	device_model.Vmp = "test"
	device_model.SchoolGuid = tool.Uuid()
	device_model.Group = 1
	device_model.Description = "出校门"
	device_model.Status = 1
	device_model.Enabled = 1
	fmt.Printf("%v\n", device_model)
	id, _ := o.Insert(device_model)
	fmt.Printf("%v\n", id)
}
