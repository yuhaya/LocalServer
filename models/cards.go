package models

import (
	"github.com/astaxie/beego/orm"
)

type Cards struct {
	Id         uint64 `orm:"fk;auto"`
	Guid       string `orm:"unique;size(50)"`
	Kind       int8   `orm:"default(1)"`
	SchoolGuid string `orm:"size(50)"`
	FamilyGuid string `orm:"size(50)"`
	Enabled    int8   `orm:"default(1)"`
}

type CardsModel struct {
}

func (this *CardsModel) List() ([]*Cards, int64, error) {
	o := orm.NewOrm()
	card_model := new(Cards)
	qs := o.QueryTable(card_model)
	var cards_list []*Cards
	num, err := qs.Filter("Enabled", 1).All(&cards_list)
	return cards_list, num, err
}

/**
 * 更具guid获取所有的卡号
 */
func (this *Cards) GetAllCardsByGuids(guids []string) map[string][]string {

}
