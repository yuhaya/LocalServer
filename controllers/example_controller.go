package controllers

type ExampleController struct {
	BaseController
}

func (this *ExampleController) Index() {
    this.TplNames = "example/index.tpl"
}
