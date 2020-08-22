package admin

import (
	"time"

	"beego-admin/utils/bcrypt"
	"beego-admin/utils/jwt"
)

type LoginController struct {
	BaseController
}

// 登录
func (c *LoginController) Login() {
	username := c.Ctx.Input.Query("username")
	if username == "" {
		c.AjaxError("请输入用户名")
	}

	adminUser := adminUserModel.GetByUsername(username)
	if adminUser.Id == 0 {
		c.AjaxError("该用户不存在")
	}

	password := c.Ctx.Input.Query("password")
	if password == "" {
		c.AjaxError("请输入密码")
	}

	dbPassword := adminUser.Password
	err := bcrypt.Compare(dbPassword, password)
	if err != nil {
		c.AjaxError("用户名或密码错误")
	}

	easyToken := jwt.EasyToken{
		Username: username,
		Expires:  time.Now().Unix() + 3600,
	}
	token, _ := easyToken.GenerateToken()

	c.AjaxSuccess(token)
}
