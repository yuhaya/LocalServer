package models

import (
	"LocalServer/tool"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type Families struct {
	Id         uint64 `orm:"fk;auto"`
	Name       string `orm:"size(50)"`
	Guid       string `orm:"unique;size(50)"`
	FirstGuid  string `orm:"size(50)"`
	SchoolGuid string `orm:"size(50)"`
}

/**
 * 更新主家长
 */
func (this *Families) UpdateMain(guid string, first_guid string) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable(this).Filter("guid", guid).Update(orm.Params{
		"first_guid": first_guid,
	})

	if err == nil {
		return true
	} else {
		return false
	}

}

type FamiliyModel struct {
	Base
}

/**
 * 根据家庭名称以检索家庭以及主家长信息
 */
func (this *FamiliyModel) GetListByFamlilyName(name string) ([]orm.Params, int64) {
	o := orm.NewOrm()
	var maps []orm.Params
	name_str := "%" + name + "%"
	sql := `SELECT f.*,u.guid as user_guid,u.phone as user_phone,u.realname as user_realname
	FROM ittr_families AS f LEFT JOIN ittr_users as u ON f.first_guid = u.guid WHERE f.name LIKE ?`
	r := o.Raw(sql, name_str)
	num, err := r.Values(&maps)
	fmt.Printf("%d==\n", num)
	if err == nil && num > 0 {
		fmt.Printf("%v===\n", maps)
	}
	return maps, num
}

/**
 *根据用户名检索家庭信息
 */
func (this *FamiliyModel) GetListByUserName(name string) ([]orm.Params, int64) {
	o := orm.NewOrm()
	var maps []orm.Params
	name_str := "%" + name + "%"
	//家长表信息检索
	sql := `SELECT f.*,u.guid as user_guid,u.phone as user_phone,u.realname as user_realname
			FROM ittr_users AS u INNER JOIN ittr_family_member AS fm ON u.guid = fm.member_guid
			INNER JOIN ittr_families as f on fm.family_guid = f.guid where u.realname LIKE ?`
	r := o.Raw(sql, name_str)
	num, err := r.Values(&maps)
	if err == nil && num > 0 {
		fmt.Printf("%v===\n", maps)
	}

	//学生信息检索
	var maps_stu []orm.Params
	sql_stu := `SELECT f.*,u.guid as user_guid,u.realname as user_realname
			FROM ittr_students AS u INNER JOIN ittr_family_member AS fm ON u.guid = fm.member_guid
			INNER JOIN ittr_families as f on fm.family_guid = f.guid where u.realname LIKE ?`
	r_stu := o.Raw(sql_stu, name_str)
	num_stu, err_stu := r_stu.Values(&maps_stu)
	if err_stu == nil && num_stu > 0 {
		fmt.Printf("%v===\n", maps_stu)
	}

	res := append(maps, maps_stu...)
	return res, num_stu + num
}

/*
 * 根据用户的手机号进行检索
 */
func (this *FamiliyModel) GetListByUserPhone(phone string) ([]orm.Params, int64) {
	o := orm.NewOrm()
	var maps []orm.Params
	is_phone := tool.IsPhone(phone)
	if !is_phone {
		return maps, 0
	}

	//家长表信息检索
	sql := `SELECT f.*,u.guid as user_guid,u.phone as user_phone,u.realname as user_realname
			FROM ittr_users AS u INNER JOIN ittr_family_member AS fm ON u.guid = fm.member_guid
			INNER JOIN ittr_families as f on fm.family_guid = f.guid where u.phone = ?`
	r := o.Raw(sql, phone)
	num, err := r.Values(&maps)
	if err == nil && num > 0 {
		fmt.Printf("%v==xxxxxxxxxxxxxxxxx=\n", maps)
	} else {
		fmt.Printf("%v====error===\n", err.Error())
	}

	return maps, num
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
func (this *FamiliyModel) AddFamily(family_model *Families, user_model *Users, family_member_model *FamilyMember) bool {

	o := orm.NewOrm()
	err := o.Begin()
	o.Insert(family_model)
	o.Insert(user_model)
	o.Insert(family_member_model)
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

/**
 * 返回所有的家庭总数
 */
func (this *FamiliyModel) CountAll() int64 {
	o := orm.NewOrm()
	sql := `SELECT COUNT(*) as total FROM ittr_families`
	r := o.Raw(sql)
	var res []orm.Params
	num, err := r.Values(&res)
	total := res[0]["total"]

	fmt.Printf("%v======\n", total)
	if err == nil && num > 0 {
		if val_str, ok := total.(string); ok {

			b, error := strconv.Atoi(val_str)
			if error == nil {
				return int64(b)
			} else {
				return 0
			}

		} else {
			return 0
		}
	} else {
		return 0
	}
}

/**
 * 依靠guid检索一个家庭对象
 */
func (this *FamiliyModel) GetFamilyByGuid(guid string) *Families {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Families))
	var fm Families
	qs.Filter("guid", guid).One(&fm)
	return &fm
}

/**
 * 更新家庭的名称
 */
func (this *FamiliyModel) UpdateFamilyNameByGuid(guid string, name string) bool {
	o := orm.NewOrm()
	num, err := o.QueryTable(new(Families)).Filter("guid", guid).Update(orm.Params{
		"name": name,
	})
	if num != 0 && err == nil {
		return true
	} else {
		return false
	}
}

type result struct {
	Id         uint64
	Guid       string
	Phone      string
	Realname   string
	Spell      string
	Gender     int8
	CreateTime time.Time
}

/**
 * 根据guid获取所有的家庭成员
 */
func (this *FamiliyModel) GetMembersByGuid(guid string) (map[string][]*result, string) {

	o := orm.NewOrm()
	var res map[string][]*result
	res = make(map[string][]*result)

	//获取所有的家长成
	qs := o.QueryTable(new(Families))
	var fm Families
	qs.Filter("guid", guid).One(&fm, "FirstGuid")
	first_guid := fm.FirstGuid

	//家长数据
	var user_maps []result
	sql := "SELECT u.* FROM ittr_family_member AS m INNER JOIN ittr_users AS u ON m.member_guid = u.guid where m.family_guid = ? AND m.type = 1"
	r := o.Raw(sql, guid)
	num, err := r.QueryRows(&user_maps)

	if err == nil && num != 0 {
		res["users"] = make([]*result, num)
		for index, _ := range user_maps {
			res["users"][index] = &user_maps[index]
		}
	} else {
		res["users"] = make([]*result, 0)
	}

	//学生数据
	var stu_maps []result
	stu_sql := `SELECT u.* FROM ittr_family_member AS m INNER JOIN ittr_students AS u ON m.member_guid = u.guid where m.family_guid = ? AND m.type = 0`
	stu_r := o.Raw(stu_sql, guid)
	stu_num, stu_err := stu_r.QueryRows(&stu_maps)

	if stu_err == nil && stu_num != 0 {

		res["stus"] = make([]*result, stu_num)
		for index, _ := range stu_maps {
			res["stus"][index] = &stu_maps[index]
		}
	} else {
		res["stus"] = make([]*result, 0)
	}

	return res, first_guid

}

/**
 * 给一个家庭添加家长
 */
func (this *FamiliyModel) AddUser(user_model *Users, family_member_model *FamilyMember) bool {

	o := orm.NewOrm()
	err := o.Begin()

	o.Insert(user_model)
	o.Insert(family_member_model)

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

/**
 * 添加一个学生
 */
func (this *FamiliyModel) AddStu(stu_model *Students, family_member_model *FamilyMember) bool {

	o := orm.NewOrm()
	err := o.Begin()

	o.Insert(stu_model)
	o.Insert(family_member_model)

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
