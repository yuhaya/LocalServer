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
	beego.Router("/card", &controllers.CardController{}, "get:Index")
	beego.Router("/card/list", &controllers.CardController{}, "get:Manager")
	beego.Router("/card/Show", &controllers.CardController{}, "get:Show")
	beego.Router("/card/Show", &controllers.CardController{}, "post:Search")
	beego.Router("/family", &controllers.FamilyController{}, "get:Index")
	beego.Router("/family/add", &controllers.FamilyController{}, "get:Add;post:AddSubmit")
}
