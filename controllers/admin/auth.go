package admin

import (
	"fmt"

	"beego-admin/models"
	"beego-admin/utils/jwt"
)

type AuthController struct {
	BaseController
	AdminUser models.AdminUser
}

func (c *AuthController) Prepare() {
	token := c.Ctx.Input.Query("token")
	if token == "" {
		c.AjaxError("您还没有登录")
	}
	// token := c.Ctx.Request.Header.Get("x-token")
	easyToken := jwt.EasyToken{}
	userName, err := easyToken.ParseToken(token)
	if err != nil {
		c.AjaxError(err.Error())
	}
	adminUser := adminUserModel.GetByUsername(userName)
	if adminUser.Id == 0 {
		c.AjaxError(fmt.Sprintf("该用户不存在: %s", userName))
	}
	c.AdminUser = adminUser
}

func (c *AuthController) GetAdminUser() {
	c.AjaxSuccess(c.AdminUser)
}
