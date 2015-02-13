package models

import (
	"time"
)

type Students struct {
	Id          uint64    `orm:"fk;auto"`
	Guid        string    `orm:"unique;size(50)"`
	Sid         string    `orm:"size(50);null"`
	Realname    string    `orm:"size(10)"`
	Spell       string    `orm:"size(10)"`
	Gender      int8      `orm:"default(1)"`
	Picture     string    `orm:"size(255);null"`
	Birthday    time.Time `orm:"default(0000-00-00);type(date);"`
	School_guid string    `orm:"size(50)"`
	Grade_guid  string    `orm:"size(50)"`
	Classname   string    `orm:"size(20);null"`
	Enrol_time  time.Time `orm:"default(0000-00-00);type(date);"`
	Create_time time.Time `orm:"type(datetime)"`
}
