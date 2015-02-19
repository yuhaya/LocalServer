package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

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
	GradeGuid string
	GradeName string
	ClassGuid string
	ClassName string
}

/**
 * 获取年纪班级数据
 */
func GetAllGradeClass() (map[string][]map[string]string, []map[string]string) {
	o := orm.NewOrm()
	sql := `SELECT g.guid as grade_guid,g.name as grade_name,c.guid as class_guid,c.name as class_name FROM
	 ittr_grades as g LEFT JOIN ittr_grade_class as gc ON gc.grade_guid = g.guid LEFT JOIN
	 ittr_classes as c  ON gc.class_guid = c.guid ORDER BY g.rating ASC`

	var res []GradeClassData
	var grade []map[string]string

	var grade_class map[string][]map[string]string
	grade_class = make(map[string][]map[string]string)

	num, _ := o.Raw(sql).QueryRows(&res)

	if num > 0 {

		tmp_grade := make(map[string]int)

		for index := int64(0); index < num; index++ {

			if tmp_grade[res[index].GradeGuid] == 0 {
				tmp_grade[res[index].GradeGuid] = 1
				grade = append(grade, map[string]string{"guid": res[index].GradeGuid, "name": res[index].GradeName})
			}

			if grade_class[res[index].GradeGuid] == nil {

				for index2 := int64(0); index2 < num; index2++ {

					if res[index2].GradeGuid == res[index].GradeGuid {

						class_guid := res[index2].ClassGuid
						class_name := res[index2].ClassName
						grade_class[res[index2].GradeGuid] = append(grade_class[res[index2].GradeGuid], map[string]string{"guid": class_guid, "name": class_name})

					}

				}
			}
		}
	}

	fmt.Printf("\n%v=======\n", grade)
	return grade_class, grade
}
