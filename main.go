package main

import (
	_ "LocalServer/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {
	orm.Debug = true
	beego.Run()
}
