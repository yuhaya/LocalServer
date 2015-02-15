package routers

import (
	"LocalServer/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Main")
	beego.Router("/frame/top", &controllers.MainController{}, "get:TopBody")
	beego.Router("/frame/left", &controllers.MainController{}, "get:LeftBody")
	beego.Router("/frame/right", &controllers.MainController{}, "get:RightBody")
	beego.Router("/example", &controllers.ExampleController{}, "get:Index")
    grade()
    class()
}

func grade(){
    beego.Router("/grade/show",&controllers.GradeController{},"get:Show")
    beego.Router("/grade",&controllers.GradeController{},"get:Index")
    beego.Router("/grade/create",&controllers.GradeController{},"post:Create")
    beego.Router("/grade/delete",&controllers.GradeController{},"post:Delete")
    beego.Router("/grade/update",&controllers.GradeController{},"post:Update")
}

func class(){
    beego.Router("/class/list",&controllers.ClassController{},"post:Index")
    beego.Router("/class/create",&controllers.ClassController{},"post:Create")
    beego.Router("/class/delete",&controllers.ClassController{},"post:Delete")
}