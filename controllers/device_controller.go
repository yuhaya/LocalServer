// @APIVersion 1.0.0
// @Title demo API
// @Description ${PROJECT_NAME} | this is demo description
// @Contact mao | 3wmaocomputer@gmail.com
// @Date 3/1/15

package controllers

import (
	"LocalServer/models"
)

type DeviceController struct {
	BaseController
}

func (this *DeviceController) Dmess() {
	de := models.DevicesModel{}
	devices, num, err := de.List()
	if num > 0 && err == nil {
		this.AjaxReturnFun("0", "success", devices)
	} else {
		this.AjaxReturnFun("-1", err.Error(), devices)
	}
}
