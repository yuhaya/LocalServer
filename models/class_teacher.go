package models

type ClassTeacher struct {
	Id          uint64 `orm:"fk;auto"`
	ClassGuid   string `orm:"size(50)"`
	TeacherGuid string `orm:"size(50)"`
	Adviser     int8   `orm:"default(0)"`
	SchoolGuid  string `orm:"size(50)"`
}
