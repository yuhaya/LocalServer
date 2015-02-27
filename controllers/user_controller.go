package controllers

type UserController struct{
    BaseController
}

func (this *UserController) Index(){
    this.TplNames="user/index.tpl"
}