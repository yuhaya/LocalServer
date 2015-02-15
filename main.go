package main

import (
	_ "LocalServer/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
)

func main() {
	beego.AddFuncMap("i18n", i18n.Tr)
	orm.Debug = true
	beego.Run()
}
