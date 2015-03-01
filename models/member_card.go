package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type MemberCard struct {
	Id         uint64 `orm:"fk;auto"`
	Card       string `orm:"size(50)"`
	Guid       string `orm:"size(50)"`
	Type       int8   `orm:"default(1)"`
	SchoolGuid string `orm:"size(50)"`
}

/**
 * 添加一张卡号
 */
func (this *MemberCard) InsertOne() (bool, int64, string) {
	o := orm.NewOrm()

	var card Cards
	cnt, _ := o.QueryTable(&card).Filter("guid", this.Card).Count()
	if cnt == 0 {
		card.SchoolGuid = this.SchoolGuid
		card.Guid = this.Card
		card.Enabled = 1
		card.Kind = 0
		_, err2 := o.Insert(&card)
		if err2 != nil {
			return false, 0, err2.Error()
		}
	}
	cnt2, _ := o.QueryTable(this).Filter("guid", this.Guid).Filter("card", this.Card).Count()
	if cnt2 > 0 {
		return false, 0, "该卡已经有人使用!"
	}

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

/**
 * 删除成员与卡的绑定关系
 */
func (this *MemberCard) Delete() bool {
	o := orm.NewOrm()
	fmt.Println(this)
	err := o.Begin()

	_, err3 := o.QueryTable(this).Filter("guid", this.Guid).Filter("card", this.Card).Delete()
	card := Cards{}
	_, err4 := o.QueryTable(&card).Filter("guid", this.Card).Delete()

	if err == nil && err3 == nil && err4 == nil {
		o.Commit()
		return true
	} else {
		o.Rollback()
		return false
	}

}

type point_result struct {
	name    string
	picture string
	grade   string
	class   string
}

/**
 * 根据卡号检索学生相关信息
 */
func (this *MemberCard) GetStuMsg() *point_result {
	sql := `SELECT s.realname as name,s.picture as picture,g.name,c.name FROM ittr_member_card AS mc INNER JOIN ittr_student AS s
	ON mc.guid = s.guid AND mc.type = 0 INNER JOIN ittr_grades
	AS g ON s.grade_guid = g.guid INNER JOIN ittr_classes AS c ON s.class_guid = c.guid WHERE mc.card = ?`
	o := orm.NewOrm()
	r := o.Raw(sql, this.Card)
	var stu_maps point_result
	r.QueryRow(&stu_maps)
	return &stu_maps
}
