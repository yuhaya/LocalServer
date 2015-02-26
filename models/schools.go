package models

import (
	"github.com/astaxie/beego/orm"
	"time"

    "strconv"
    "fmt")

type Schools struct {
	Id         uint64    `orm:"fk;auto"`
	Guid       string    `orm:"unique;size(50)"`
	Name       string    `orm:"size(30)"`
	Spell      string    `orm:"size(50)"`
	Province   string    `orm:"size(20)"`
	City       string    `orm:"size(20)"`
	County     string    `orm:"size(20)"`
	Address    string    `orm:"size(80);null"`
	UpdateTime time.Time `orm:"type(datetime)"`
	DoorNum    int8    `orm:"default(1)"`
	Enabled    int8      `orm:"default(1)"`
}
type SchoolModel struct {
}

//获取学校guid
func (this *SchoolModel) GetSchoolGuid() string {
	var school Schools
	err := orm.NewOrm().QueryTable(new(Schools)).One(&school)
	if err == orm.ErrMultiRows {
		return ""
	}
	if err == orm.ErrNoRows {
		return ""
	}
	return school.Guid
}

func (this *Schools) Insert() error{
    if _,err:=orm.NewOrm().Insert(this);err!=nil{
        return err
    }
    return nil
}

func (this *Schools) Update(fields ...string) error{
    var sql ="UPDATE `ittr_schools` set"+
    " `name`='"+this.Name+"'"+
    ", `province`='"+this.Province+"'"+
    ", `city`='"+this.City+"'"+
    ", `county`='"+this.County+"'"+
    ", `address`='"+this.Address+"'"+
    ", `door_num`="+strconv.Itoa(int(this.DoorNum))+
    " where `guid`='"+this.Guid+"'";
    o:=orm.NewOrm()
    fmt.Println(sql)
    _,err:=o.Raw(sql).Exec()
    if err !=nil{
        return err
    }
    return nil
}

func (this *Schools) Query(guid string) orm.QuerySeter{
    return orm.NewOrm().QueryTable(this).Filter("guid",guid)
}
