package models
import "github.com/astaxie/beego/orm"

type GradeClass struct {
	Id         uint64 `orm:"fk;auto"`
	GradeGuid  string `orm:"size(50)"`
	ClassGuid  string `orm:"size(50)"`
	SchoolGuid string `orm:"size(50)"`
}

func (g *GradeClass) Insert() error{
    if _,err:=orm.NewOrm().Insert(g);err!=nil{
        return err
    }
    return nil
}

func (g *GradeClass) InsertTrans(o orm.Ormer) error{
    if _,err:=o.Insert(g);err!=nil{
        return err
    }
    return nil
}

func (g *GradeClass) DeleteTrans(guid string,class_guid string,o orm.Ormer) error{
    if _,err:=o.QueryTable(g).Filter("grade_guid",guid).Filter("class_guid",class_guid).Delete();err!=nil{
        return err
    }
    return nil
}