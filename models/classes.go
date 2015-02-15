package models
import "github.com/astaxie/beego/orm"

type Classes struct {
    Id         uint64 `orm:"fk;auto"`
    Guid       string `orm:"unique;size(50)"`
    Name       string `orm:"size(20)"`
    SchoolGuid string `orm:"size(50)"`
}

type Class struct {
    Guid string
    GradeGuid string
    Name string
    Teacher string
    Adviser uint64
    TeacherGuid string
}

func (c *Classes) Query(guid string) (class []Class, err error) {

    var classes []Class
    var sql = "SELECT ittr_teachers.name AS teacher, "+
    " ittr_class_teacher.adviser AS adviser,"+
    " ittr_classes.name AS name,"+
    " ittr_classes.guid AS guid,"+
    " ittr_grade_class.grade_guid AS grade_guid,"+
    " ittr_teachers.guid AS teacher_guid"+
    " FROM ittr_grade_class INNER JOIN ittr_classes ON ittr_grade_class.class_guid = ittr_classes.guid"+
    " INNER JOIN ittr_class_teacher ON ittr_classes.guid = ittr_class_teacher.class_guid"+
    " INNER JOIN ittr_teachers ON ittr_class_teacher.teacher_guid = ittr_teachers.guid"+
    " WHERE ittr_grade_class.grade_guid= ? ";
    o := orm.NewOrm()
    var r orm.RawSeter
    r = o.Raw(sql, guid)
    num, err := r.QueryRows(&classes)
    if err!=nil {
        return nil, err
    }
    if num>0 {
        return classes, nil
    }
    return nil, nil
}

func (c *Classes) Insert() error{
    if _,err:=orm.NewOrm().Insert(c);err!=nil{
        return err
    }
    return nil
}

func (c *Classes) InsertTrans(o orm.Ormer) error{
    if _,err:=o.Insert(c);err!=nil{
        return err
    }
    return nil
}

func (c *Classes) DeleteTrans(guid string,o orm.Ormer) error{
    if _,err:=o.QueryTable(c).Filter("guid",guid).Delete();err!=nil{
        return err
    }
    return nil
}