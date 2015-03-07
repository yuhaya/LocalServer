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
	beego.Router("/card/add", &controllers.CardController{}, "post:Add")
	beego.Router("/card/del", &controllers.CardController{}, "*:Del")
	beego.Router("/card/members", &controllers.CardController{}, "get:MemeberList")

	beego.Router("/family", &controllers.FamilyController{}, "get:Index")
	beego.Router("/family/add", &controllers.FamilyController{}, "get:Add;post:AddSubmit")
	beego.Router("/family/edit", &controllers.FamilyController{}, "get:EditFamily;post:EditSubmit")
	beego.Router("/family/members", &controllers.FamilyController{}, "get:ManagerMember")
	beego.Router("/family/delete", &controllers.FamilyController{}, "*:Delete")
	beego.Router("/family/notice", &controllers.FamilyController{}, "*:Notice")
	beego.Router("/family/notice/bind", &controllers.FamilyController{}, "*:NoticeBind")
	beego.Router("/family/notice/delete", &controllers.FamilyController{}, "*:NoticeDeleteBind")

	beego.Router("/family/members/user", &controllers.FamilyController{}, "get:AddMemberShow;post:AddMember")
	beego.Router("/family/members/parent", &controllers.FamilyController{}, "get:ShowUser")
	beego.Router("/family/members/parent/post", &controllers.FamilyController{}, "get:EditUserShow")
	beego.Router("/family/members/parent/delete", &controllers.FamilyController{}, "*:UDelete")
	beego.Router("/family/members/stu", &controllers.FamilyController{}, "get:ShowStu")
	beego.Router("/family/members/stu/post", &controllers.FamilyController{}, "get:EditStuShow")
	beego.Router("/family/members/stu/delete", &controllers.FamilyController{}, "*:SDelete")
	beego.Router("/family/members/main", &controllers.FamilyController{}, "*:SetMainUser")
	beego.Router("/family/members/list", &controllers.FamilyController{}, "get:MemeberList")
	beego.Router("/family/card/del", &controllers.FamilyController{}, "*:Del")

	beego.Router("/point/create", &controllers.PointController{}, "post:Create")
	beego.Router("/point", &controllers.PointController{}, "*:CreateShow")

	beego.Router("/device/get", &controllers.DeviceController{}, "*:Dmess")

	beego.Router("/mode", &controllers.MainController{}, "*:Mode")
	beego.Router("/regcode", &controllers.MainController{}, "*:Card")
	beego.Router("/mode/post", &controllers.MainController{}, "post:SetMode")
	grade()
	class()
	school()
	user()
}

func grade() {
	beego.Router("/grade/show", &controllers.GradeController{}, "get:Show")
	beego.Router("/grade", &controllers.GradeController{}, "get:Index")
	beego.Router("/grade/create", &controllers.GradeController{}, "post:Create")
	beego.Router("/grade/delete", &controllers.GradeController{}, "post:Delete")
	beego.Router("/grade/update", &controllers.GradeController{}, "post:Update")
}

func class() {
	beego.Router("/class/list", &controllers.ClassController{}, "post:Index")
	beego.Router("/class/create", &controllers.ClassController{}, "post:Create")
	beego.Router("/class/delete", &controllers.ClassController{}, "post:Delete")
}

func school() {
	beego.Router("/school", &controllers.SchoolController{}, "get:Index")
	beego.Router("/school/create", &controllers.SchoolController{}, "post:Create")
}

func user() {
	beego.Router("/user/index", &controllers.UserController{}, "get:Index")
}
