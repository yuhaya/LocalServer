package models

import (
	"time"
)

type Attendances struct {
	Id         uint64    `orm:"fk;auto"`
	Card       string    `orm:"size(50)"`
	Time       time.Time `orm:"type(datetime)"`
	Type       byte      `orm:"default(0)"`
	SchoolGuid string    `orm:"size(50)"`
	Auto       byte      `orm:"default(1)"`
}
