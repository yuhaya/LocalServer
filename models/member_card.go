package models

import (
	"github.com/astaxie/beego/orm"
)

type MemberCard struct {
	Id         uint64 `orm:"fk;auto"`
	Card       string `orm:"size(50)"`
	Guid       string `orm:"size(50)"`
	Type       int8   `orm:"default(1)"`
	SchoolGuid string `orm:"size(50)"`
}

func (this *MemberCard) InsertOne() (bool, int64, string) {
	o := orm.NewOrm()
	id, err := o.Insert(this)
	if err == nil {
		return true, id, ""
	} else {
		return false, 0, err.Error()
	}
}

/**
 * 更具guid获取所有的卡号
 */
func (this *MemberCard) GetAllCardsByGuids(guids []string) map[string][]string {
	var mc []*MemberCard
	var res map[string][]string

	res = make(map[string][]string)

	o := orm.NewOrm()
	num, err := o.QueryTable(this).Filter("guid__in", guids).All(&mc)
	if num > 0 && err == nil {
		for _, record := range mc {
			val, ok := res[record.Guid]
			if ok == false {
				tmp_slice := make([]string, 0, 20)
				res[record.Guid] = append(tmp_slice, record.Card)
			} else {
				res[record.Guid] = append(val, record.Card)
			}
		}
		return res
	} else {
		return res
	}
}
