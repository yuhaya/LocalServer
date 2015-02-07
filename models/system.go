package models

type System struct {
	Id         uint64 `orm:"fk;auto"`
	Key        string `orm:"size(50)"`
	Value      string `orm:"size(250)"`
	SchoolGuid string `orm:"size(50)"`
}
