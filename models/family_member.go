package models

type FamilyMember struct {
	Id         uint64 `orm:"fk;auto"`
	FamilyGuid string `orm:"size(50)"`
	MemberGuid string `orm:"size(50)"`
	Type       byte   `orm:"default(1)"`
	SchoolGuid string `orm:"size(50)"`
}
