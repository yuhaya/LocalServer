package models
import "github.com/astaxie/beego/orm"

type Teachers struct {
	Id         uint64 `orm:"fk;auto"`
    Guid       string `orm:"size(50)"`
	Name       string `orm:"size(50)"`
	Phone      string `orm:"size(20)"`
	Gender     byte   `orm:"default(0)"`
	SchoolGuid string `orm:"size(50)"`
}

func (t *Teachers) Insert() error{
    if _,err:=orm.NewOrm().Insert(t);err!=nil{
        return err
    }
    return nil
}

func (t *Teachers) InsertTrans(o orm.Ormer) error{
    if _,err:=o.Insert(t);err!=nil{
        return err
    }
    return nil
}

func (t *Teachers) DeleteTrans(guid string,o orm.Ormer) error{
    if _,err:=o.QueryTable(t).Filter("guid",guid).Delete();err!=nil{
        return err
    }
    return nil
}