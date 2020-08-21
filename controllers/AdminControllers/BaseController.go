package AdminControllers

import (
	"beego-admin/controllers"
	"beego-admin/helpers"
)

type BaseController struct {
	controllers.BaseController
}

func (c *BaseController) Prepare() {
	token := c.Ctx.Input.Query("token")
	if token == "" {
		c.AjaxError("您还没有登录")
	}
	// token := c.Ctx.Request.Header.Get("x-token")
	easyToken := helpers.EasyToken{}
	_, err := easyToken.ParseToken(token)
	if err != nil {
		c.AjaxError(err.Error())
	}
}
