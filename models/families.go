package models

import (
	"LocalServer/tool"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Families struct {
	Id         uint64 `orm:"fk;auto"`
	Name       string `orm:"size(50)"`
	Guid       string `orm:"unique;size(50)"`
	FirstGuid  string `orm:"size(50)"`
	SchoolGuid string `orm:"size(50)"`
}

type FamiliyModel struct {
}

/**
 * 根据家庭名称以检索家庭以及主家长信息
 */
func (this *FamiliyModel) GetListByFamlilyName(name string) []orm.Params {
	o := orm.NewOrm()
	var maps []orm.Params
	name_str := "%" + name + "%"
	sql := `SELECT f.*,u.guid as user_guid,u.phone as user_guid,u.realname as user_realname
	FROM ittr_families AS f LEFT JOIN ittr_users as u ON f.first_guid = u.guid WHERE f.name LIKE ?`
	r := o.Raw(sql, name_str)
	num, err := r.Values(&maps)
	fmt.Printf("%d==\n", num)
	if err == nil && num > 0 {
		fmt.Printf("%v===\n", maps)
	}
	return maps
}

/**
 *根据用户名检索家庭信息
 */
func (this *FamiliyModel) GetListByUserName(name string) []orm.Params {
	o := orm.NewOrm()
	var maps []orm.Params
	name_str := "%" + name + "%"
	//家长表信息检索
	sql := `SELECT f.*,u.guid as user_guid,u.phone as user_phone,u.realname as user_realname
			FROM ittr_users AS u INNER JOIN ittr_family_member AS fm ON u.guid = fm.member_guid
			INNER JOIN ittr_families as f on fm.family_guid = f.guid where f.name LIKE ?`
	r := o.Raw(sql, name_str)
	num, err := r.Values(&maps)
	if err == nil && num > 0 {
		fmt.Printf("%v===\n", maps)
	}

	//学生信息检索
	var maps_stu []orm.Params
	sql_stu := `SELECT f.*,u.guid as user_guid,u.phone as user_phone,u.realname as user_realname
			FROM ittr_students AS u INNER JOIN ittr_family_member AS fm ON u.guid = fm.member_guid
			INNER JOIN ittr_families as f on fm.family_guid = f.guid where f.name LIKE ?`
	r_stu := o.Raw(sql_stu, name_str)
	num_stu, err_stu := r_stu.Values(&maps_stu)
	if err_stu == nil && num_stu > 0 {
		fmt.Printf("%v===\n", maps_stu)
	}

	res := append(maps, maps_stu...)
	return res
}

/*
 * 根据用户的手机号进行检索
 */
func (this *FamiliyModel) GetListByUserPhone(phone string) []orm.Params {
	o := orm.NewOrm()
	var maps []orm.Params
	is_phone := tool.IsPhone(phone)
	if !is_phone {
		return maps
	}

	//家长表信息检索
	sql := `SELECT f.*,u.guid as user_guid,u.phone as user_phone,u.realname as user_realname
			FROM ittr_users AS u INNER JOIN ittr_family_member AS fm ON u.guid = fm.member_guid
			INNER JOIN ittr_families as f on fm.family_guid = f.guid where f.phone = ?`
	r := o.Raw(sql, phone)
	num, err := r.Values(&maps)
	if err == nil && num > 0 {
		fmt.Printf("%v===\n", maps)
	}

	return maps
}

/**
 * 获取所有的家庭信息
 */
func (this *FamiliyModel) GetAllPage(offset int, page int) ([]orm.Params, int64) {
	o := orm.NewOrm()
	var maps []orm.Params
	sql := `SELECT f.*,u.guid as user_guid,u.phone as user_phone,u.realname as user_realname
	FROM ittr_families AS f LEFT JOIN ittr_users as u ON f.first_guid = u.guid LIMIT ?,?`
	r := o.Raw(sql, offset, page)
	num, err := r.Values(&maps)
	if err == nil && num > 0 {
		fmt.Printf("%v===\n", maps)
	}
	return maps, num
}

/**
 * 添加一个家庭
 */
func (this *FamiliyModel) AddFamily(family_model *Families, user_model *Users) bool {

	o := orm.NewOrm()
	err := o.Begin()
	o.Insert(family_model)
	o.Insert(user_model)
	var err2 error
	if err == nil {
		err2 = o.Commit()
	} else {
		err2 = o.Rollback()
	}
	if err == nil && err2 == nil {
		return true
	} else {
		return false
	}
}
