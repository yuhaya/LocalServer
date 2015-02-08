package models

type GradeClass struct {
	Id         uint64 `orm:"fk;auto"`
	GradeGuid  string `orm:"size(50)"`
	ClassGuid  string `orm:"size(50)"`
	SchoolGuid string `orm:"size(50)"`
}
