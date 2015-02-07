package models

type Devices struct {
	Id          uint64 `orm:"fk;auto"`
	Guid        string `orm:"unique;size(50)"`
	Device      string `orm:"size(50)"`
	Kind        byte   `orm:"default(0)"`
	Vmp         string `orm:"size(10)"`
	SchoolGuid  string `orm:"size(50)"`
	Group       byte   `orm:"default(0)"`
	Description string `orm:"size(255)"`
	Enabled     byte   `orm:"default(1)"`
}
