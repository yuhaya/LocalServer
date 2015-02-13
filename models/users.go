package models

import (
	"time"
)

type Users struct {
	Id         uint64    `orm:"fk;auto"`
	Guid       string    `orm:"unique;size(50)"`
	Phone      string    `orm:"size(20);null"`
	Realname   string    `orm:"size(10);null"`
	Password   string    `orm:"size(50)"`
	Spell      string    `orm:"size(10)"`
	Gender     int8      `orm:"default(1)"`
	IdCard     string    `orm:"size(20);null"`
	Picture    string    `orm:"size(100);null"`
	SchoolGuid string    `orm:"size(50)"`
	CreateTime time.Time `orm:"type(datetime)"`
}
