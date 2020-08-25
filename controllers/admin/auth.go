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
	token := c.Ctx.Request.Header.Get("AUTHORIZATION")
	if token == "" {
		c.ErrorJson("您还没有登录")
	}
	easyToken := jwt.EasyToken{}
	userName, err := easyToken.ParseToken(token)
	if err != nil {
		c.ErrorJson(err.Error())
	}
	adminUser := adminUserModel.GetByUsername(userName)
	if adminUser.Id == 0 {
		c.ErrorJson(fmt.Sprintf("该用户不存在: %s", userName))
	}
	c.AdminUser = adminUser
}

func (c *AuthController) GetAdminUser() {
	c.SuccessJson(c.AdminUser)
}
