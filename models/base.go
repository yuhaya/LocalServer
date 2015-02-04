package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"time"
)

type Base struct {
}

type Students struct {
	Id          uint64    `orm:"fk;auto"`
	Guid        string    `orm:"unique;size(50)"`
	Sid         string    `orm:"size(50);null"`
	Realname    string    `orm:"size(10)"`
	Spell       string    `orm:"size(10)"`
	Gender      byte      `orm:"default(1)"`
	Picture     string    `orm:"size(255);null"`
	Birthday    time.Time `orm:"default(0000-00-00);type(date);"`
	School_guid string    `orm:"size(50)"`
	Grade_guid  string    `orm:"size(50)"`
	Classname   string    `orm:"size(20);null"`
	Enrol_time  time.Time `orm:"default(0000-00-00);type(date);"`
	Create_time time.Time `orm:"type(datetime)"`
}

type Users struct {
	Id         uint64    `orm:"fk;auto"`
	Guid       string    `orm:"unique;size(50)"`
	Phone      string    `orm:"size(20);null"`
	Realname   string    `orm:"size(10);null"`
	Password   string    `orm:"size(50)"`
	Spell      string    `orm:"size(10)"`
	Gender     byte      `orm:"default(1)"`
	IdCard     string    `orm:"size(20);null"`
	Picture    string    `orm:"size(100);null"`
	SchoolGuid string    `orm:"size(50)"`
	CreateTime time.Time `orm:"type(datetime)"`
}

type Cards struct {
	Id         uint64 `orm:"fk;auto"`
	Guid       string `orm:"unique;size(50)"`
	Kind       byte   `orm:"default(1)"`
	SchoolGuid string `orm:"size(50)"`
	FamilyGuid string `orm:"size(50)"`
	Enabled    byte   `orm:"default(1)"`
}

type Families struct {
	Id         uint64 `orm:"fk;auto"`
	Guid       string `orm:"unique;size(50)"`
	FirstGuid  string `orm:"size(50)"`
	SchoolGuid string `orm:"size(50)"`
}

type FamilyMember struct {
	Id         uint64 `orm:"fk;auto"`
	FamilyGuid string `orm:"size(50)"`
	MemberGuid string `orm:"size(50)"`
	Type       byte   `orm:"default(1)"`
	SchoolGuid string `orm:"size(50)"`
}
type FamilyRelation struct {
	Id         uint64 `orm:"fk;auto"`
	UserGuid   string `orm:"size(50)"`
	StuGuid    string `orm:"size(50)"`
	Relation   string `orm:"size(50)"`
	SchoolGuid string `orm:"size(50)"`
}

type MemberCard struct {
	Id         uint64 `orm:"fk;auto"`
	Card       string `orm:"size(50)"`
	Guid       string `orm:"size(50)"`
	Type       byte   `orm:"default(1)"`
	SchoolGuid string `orm:"size(50)"`
}

type CardReceiver struct {
	Id         uint64 `orm:"fk;auto"`
	Card       string `orm:"size(50)"`
	Guid       string `orm:"size(50)"`
	Type       byte   `orm:"default(1)"`
	SchoolGuid string `orm:"size(50)"`
}

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
	DoorNum    int16     `orm:"default(1)"`
	Enabled    byte      `orm:"default(1)"`
}

type Devices struct {
	Id          uint64 `orm:"fk;auto"`
	Guid        string `orm:"unique;size(50)"`
	Device      string `orm:"size(50)"`
	Kind        byte   `orm:"default(0)"`
	Vmp         string `orm:"size(10)"`
	SchoolGuid  string `orm:"size(50)"`
	Group       byte   `orm:"default(0)"`
	Description string `orm:"size(255)"`
	Enabled     byte   `orm:"default(1)"`
}

type Admins struct {
	Id         uint64    `orm:"fk;auto"`
	Guid       string    `orm:"unique;size(50)"`
	Username   string    `orm:"size(50)"`
	Password   string    `orm:"size(50)"`
	Realname   string    `orm:"size(50);null"`
	SchoolGuid string    `orm:"size(50)"`
	CreateTime time.Time `orm:"type(datetime)"`
	Super      byte      `orm:"default(0)"`
	Enabled    byte      `orm:"default(1)"`
}

type Attendances struct {
	Id         uint64    `orm:"fk;auto"`
	Card       string    `orm:"size(50)"`
	Time       time.Time `orm:"type(datetime)"`
	Type       byte      `orm:"default(0)"`
	SchoolGuid string    `orm:"size(50)"`
	Auto       byte      `orm:"default(1)"`
}

type System struct {
	Id         uint64 `orm:"fk;auto"`
	Key        string `orm:"size(50)"`
	Value      string `orm:"size(250)"`
	SchoolGuid string `orm:"size(50)"`
}

type Grades struct {
	Id          uint64 `orm:"fk;auto"`
	Guid        string `orm:"size(50)"`
	Name        string `orm:"size(50)"`
	Rating      byte   `orm:"default(0)"`
	ClassNumber byte   `orm:"default(1)"`
	SchoolGuid  string `orm:"size(50)"`
}

func Regist() {
	dbUser := beego.AppConfig.String("dbUser")
	dbName := beego.AppConfig.String("dbName") //数据库名字
	dbPwd := beego.AppConfig.String("dbPwd")
	dbLink := fmt.Sprintf("%s:%s@/%s?charset=utf8", dbUser, dbPwd, dbName) //数据库连接字符串
	orm.RegisterModelWithPrefix("ittr_",
		new(Students),
		new(Admins),
		new(Attendances),
		new(CardReceiver),
		new(Cards),
		new(Devices),
		new(Families),
		new(FamilyMember),
		new(FamilyRelation),
		new(Grades),
		new(MemberCard),
		new(Schools),
		new(System),
		new(Users))
	orm.RegisterDriver("mysql", orm.DR_MySQL)        //注册数据库驱动
	orm.RegisterDataBase("default", "mysql", dbLink) //注册数据库，并设置默认数据库
}

func init() {
	Regist()
	createTable()
}

func createTable() {
	name := "default"                          //数据库别名
	force := true                              //不强制建数据库
	verbose := true                            //打印建表过程
	err := orm.RunSyncdb(name, force, verbose) //建表
	if err != nil {
		beego.Error(err)
	}
}

func (this *Base) Demo() {
	fmt.Println("hello world\n")
}
