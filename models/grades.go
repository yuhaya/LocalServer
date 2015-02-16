package models
import "github.com/astaxie/beego/orm"

type Grades struct {
	Id          uint64 `orm:"fk;auto"`
	Guid        string `orm:"unique;size(50)"`
	Name        string `orm:"size(50)"`
	Rating      uint64   `orm:"default(0)"`
	SchoolGuid  string `orm:"size(50)"`
}

func (m *Grades) Insert() error{
    if _,err:=orm.NewOrm().Insert(m);err!=nil{
        return err
    }
    return nil
}

func (m *Grades) Delete(guid string) error{
    if _,err:=orm.NewOrm().QueryTable(m).Filter("guid",guid).Delete();err!=nil{
        return err
    }
    return nil
}

func (m *Grades) Update(fields ...string) error{
    if _,err := orm.NewOrm().Update(m,fields...);err!=nil{
        return err
    }
    return nil
}

func (m *Grades) Query() orm.QuerySeter{
    return orm.NewOrm().QueryTable(m)
}

func (m *Grades) MaxRating() uint64{
    var list []*Grades
    if count,_:=m.Query().Count();count>0{
        m.Query().OrderBy("-rating").All(&list)
    }else{
        return 0
    }
    return list[0].Rating
}