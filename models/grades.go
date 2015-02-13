package models

type Grades struct {
	Id          uint64 `orm:"fk;auto"`
	Guid        string `orm:"unique;size(50)"`
	Name        string `orm:"size(50)"`
	Rating      int8   `orm:"default(0)"`
	ClassNumber int8   `orm:"default(1)"`
	SchoolGuid  string `orm:"size(50)"`
}
