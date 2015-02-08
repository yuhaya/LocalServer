package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

type Base struct {
}

func Regist() {
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
		new(Users),
		new(ClassTeacher),
		new(Teachers),
		new(Classes),
		new(GradeClass),
	)
	orm.RegisterDriver("mysql", orm.DR_Sqlite) //注册数据库驱动
	orm.RegisterDataBase("default", "sqlite3", "data.db")
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
	fmt.Println("hello world \n")
}
