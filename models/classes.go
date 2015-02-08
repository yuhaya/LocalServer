package models

type Classes struct {
	Id         uint64 `orm:"fk;auto"`
	guid       string `orm:"unique;size(50)"`
	name       string `orm:"size(20)"`
	SchoolGuid string `orm:"size(50)"`
}
