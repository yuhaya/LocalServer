package models

type FamilyRelation struct {
	Id         uint64 `orm:"fk;auto"`
	UserGuid   string `orm:"size(50)"`
	StuGuid    string `orm:"size(50)"`
	Relation   string `orm:"size(50)"`
	SchoolGuid string `orm:"size(50)"`
}
