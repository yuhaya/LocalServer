package controllers

import (
	// "encoding/json"
	"github.com/astaxie/beego"
	"regexp"
	"strings"
	"text/template"
	"time"
)

const PAGE_NUM = 10

type AjaxReturn struct {
	Code bool
	Msg  string
	Data interface{}
}

type NestPreparer interface {
	NestPrepare()
}

type BaseController struct {
	beego.Controller
}

// Prepare implemented Prepare method for baseRouter.
func (this *BaseController) Prepare() {

	// page start time
	this.Data["PageStartTime"] = time.Now()

	// Setting properties.
	this.Data["AppDescription"] = beego.AppConfig.String("AppDescription")
	this.Data["AppKeywords"] = beego.AppConfig.String("AppKeywords")
	this.Data["AppName"] = beego.AppConfig.String("AppName")
	this.Data["AppVer"] = beego.AppConfig.String("AppVer")
	this.Data["AppUrl"] = beego.AppConfig.String("AppUrl")
	this.Data["AppLogo"] = beego.AppConfig.String("AppLogo")
	this.Data["IsProMode"] = beego.AppConfig.String("IsProMode")
	this.Data["TitleName"] = ""

	controllerName, methodName := this.GetControllerAndAction()
	reg := regexp.MustCompile(`Controller`)
	controllerName = reg.ReplaceAllString(controllerName, "")
	this.Data["ControllerName"] = strings.ToLower(controllerName)
	this.Data["MethodName"] = strings.ToLower(methodName)

	if controllerName != "Main" && !this.IsAjax() {
		this.Layout = "main/layout.tpl"
	}

	if app, ok := this.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

func (this *BaseController) AjaxReturnFun(code bool, msg string, data interface{}) {
	m := AjaxReturn{code, msg, data}
	this.Data["json"] = &m
	this.ServeJson()
	// b, err := json.Marshal(m)
	// if err == nil {
	// 	this.Ctx.WriteString(string(b))
	// } else {
	// 	this.Ctx.WriteString("{\"code\":0,\"msg\":\"系统异常\",\"data\":\"\"}")
	// }
}

func (this *BaseController) OutputMsg(msg string, urlmsg map[string]string) {
	t, _ := template.New("sysmsg.tpl").ParseFiles(beego.ViewsPath + "/sysmsg.tpl")
	data := make(map[string]interface{})
	data["content"] = msg
	data["urlmsg"] = urlmsg
	t.Execute(this.Ctx.ResponseWriter, data)
}
