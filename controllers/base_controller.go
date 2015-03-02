package controllers

import (
	// "encoding/json"
	//	"fmt"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/beego/i18n"
	"regexp"
	"strings"
	"text/template"
	"time"
)

const PAGE_NUM = 10

// langType represents a language type.
type langType struct {
	Lang, Name string
}

//var langTypes []*langType // Languages are supported.

type AjaxReturn struct {
	Code string
	Msg  string
	Data interface{}
}

type NestPreparer interface {
	NestPrepare()
}

type BaseController struct {
	beego.Controller
	i18n.Locale
	CacheObj cache.Cache
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
	this.setLangVer()

	controllerName, methodName := this.GetControllerAndAction()
	reg := regexp.MustCompile(`Controller`)
	controllerName = reg.ReplaceAllString(controllerName, "")
	this.Data["ControllerName"] = strings.ToLower(controllerName)
	this.Data["MethodName"] = strings.ToLower(methodName)

	bm, err := cache.NewCache("memory", `{"interval":0}`)
	if err != nil {
		//缓存系统故障
	} else {
		this.CacheObj = bm
	}

	if controllerName != "Main" && !this.IsAjax() {
		this.Layout = "main/layout.tpl"
	}

	if app, ok := this.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}

}

func (this *BaseController) AjaxReturnFun(code string, msg string, data interface{}) {
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

// setLangVer sets site language version.
func (this *BaseController) setLangVer() bool {
	isNeedRedir := false
	hasCookie := false

	// 1. Check URL arguments.
	lang := this.Input().Get("lang")

	// 2. Get language information from cookies.
	if len(lang) == 0 {
		lang = this.Ctx.GetCookie("lang")
		hasCookie = true
	} else {
		isNeedRedir = true
	}

	// Check again in case someone modify by purpose.
	if !i18n.IsExist(lang) {
		lang = ""
		isNeedRedir = false
		hasCookie = false
	}

	// 3. Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		al := this.Ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			al = al[:5] // Only compare first 5 letters.
			if i18n.IsExist(al) {
				lang = al
			}
		}
	}

	// 4. Default language is English.
	if len(lang) == 0 {
		lang = "zh-CN"
		isNeedRedir = false
	}

	curLang := langType{
		Lang: lang,
	}

	// Save language information in cookies.
	if !hasCookie {
		this.Ctx.SetCookie("lang", curLang.Lang, 1<<31-1, "/")
	}

	//	restLangs := make([]*langType, 0, len(langTypes)-1)
	//	for _, v := range langTypes {
	//		if lang != v.Lang {
	//			restLangs = append(restLangs, v)
	//		} else {
	//			curLang.Name = v.Name
	//		}
	//	}

	// Set language properties.
	this.Lang = lang
	this.Data["Lang"] = curLang.Lang

	this.Data["CurLang"] = curLang.Name
	//	this.Data["RestLangs"] = restLangs
	return isNeedRedir
}

func (this *BaseController) Mode() int {
	mode := 0
	if this.CacheObj.IsExist("RUNMODE") {
		mode_val := this.CacheObj.Get("RUNMODE")
		mode_ass_val, ok := mode_val.(int)
		if ok {
			mode = mode_ass_val
		}
	} else {
		this.CacheObj.Put("RUNMODE", mode, 0)
	}
	return mode
}

func (this *BaseController) SetMode() error {
	mode, _ := this.GetInt("mode")
	fmt.Printf("\n+++++++++%d++++++++++\n", mode)
	err := this.CacheObj.Put("RUNMODE", mode, 90000000000)
	return err

}
