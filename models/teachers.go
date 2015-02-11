package models

type Teachers struct {
	Id         uint64 `orm:"fk;auto"`
	name       string `orm:"size(50)"`
	phone      string `orm:"size(20)"`
	gender     int8   `orm:"default(0)"`
	SchoolGuid string `orm:"size(50)"`
}
