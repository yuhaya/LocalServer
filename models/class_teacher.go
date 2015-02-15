package models
import "github.com/astaxie/beego/orm"

type ClassTeacher struct {
	Id          uint64 `orm:"fk;auto"`
	ClassGuid   string `orm:"size(50)"`
	TeacherGuid string `orm:"size(50)"`
	Adviser     byte   `orm:"default(0)"`
	SchoolGuid  string `orm:"size(50)"`
}

func (c *ClassTeacher) Insert() error{
    if _,err:=orm.NewOrm().Insert(c);err!=nil{
        return err
    }
    return nil
}

func (c *ClassTeacher) InsertTrans(o orm.Ormer) error{
    if _,err:=o.Insert(c);err!=nil{
        return err
    }
    return nil
}

func (c *ClassTeacher) DeleteTrans(class_guid string,o orm.Ormer) error{
    if _,err:=o.QueryTable(c).Filter("class_guid",class_guid).Delete();err!=nil{
        return err
    }
    return nil
}