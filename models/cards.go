package models

type Cards struct {
	Id         uint64 `orm:"fk;auto"`
	Guid       string `orm:"unique;size(50)"`
	Kind       byte   `orm:"default(1)"`
	SchoolGuid string `orm:"size(50)"`
	FamilyGuid string `orm:"size(50)"`
	Enabled    byte   `orm:"default(1)"`
}
