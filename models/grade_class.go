package models

import "github.com/astaxie/beego/orm"

type GradeClass struct {
	Id         uint64 `orm:"fk;auto"`
	GradeGuid  string `orm:"size(50)"`
	ClassGuid  string `orm:"size(50)"`
	SchoolGuid string `orm:"size(50)"`
}

func (g *GradeClass) Insert() error {
	if _, err := orm.NewOrm().Insert(g); err != nil {
		return err
	}
	return nil
}

func (g *GradeClass) InsertTrans(o orm.Ormer) error {
	if _, err := o.Insert(g); err != nil {
		return err
	}
	return nil
}

func (g *GradeClass) DeleteTrans(guid string, class_guid string, o orm.Ormer) error {
	if _, err := o.QueryTable(g).Filter("grade_guid", guid).Filter("class_guid", class_guid).Delete(); err != nil {
		return err
	}
	return nil
}

type GradeClassData struct {
	grade_guid string
	grade_name string
	class_guid string
	class_name string
}

func GetAllGradeClass() {
	o := orm.NewOrm()
	sql := `SELECT g.guid as grade_guid,g.name as grade_name,c.guid as class_guid,c.name as class_name FROM ittr_grade_class as gc INNER JOIN ittr_grades as g ON gc.grade_guid = g.guid LEFT JOIN ittr_classes as c  ON gc.class_guid = c.guid`
	var res []GradeClassData
	var grade map[string]string
	num, _ := o.Raw(sql).QueryRows(&res)
	if num > 0 {
		for index := 2; index < num; index++ {
			if grade[res[index].grade_guid] == "" {

			}
		}
	}
}
