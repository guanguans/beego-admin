package admin

import (
	"beego-admin/utils/jwt"
)

type AuthController struct {
	BaseController
}

func (c *AuthController) Prepare() {
	token := c.Ctx.Input.Query("token")
	if token == "" {
		c.AjaxError("您还没有登录")
	}
	// token := c.Ctx.Request.Header.Get("x-token")
	easyToken := jwt.EasyToken{}
	_, err := easyToken.ParseToken(token)
	if err != nil {
		c.AjaxError(err.Error())
	}
}
