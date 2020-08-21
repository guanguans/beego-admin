package AdminControllers

import (
	"time"

	"beego-admin/helpers"
)

type AuthController struct {
	BaseController
}

// 登录
func (c *AuthController) Login() {
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
	err := helpers.Compare(dbPassword, password)
	if err != nil {
		c.AjaxError("用户名或密码错误")
	}

	easyToken := helpers.EasyToken{
		Username: username,
		Expires:  time.Now().Unix() + 3600,
	}
	token, _ := easyToken.GenerateToken()

	c.AjaxSuccess(token)
}

func (c *AuthController) GetToken() {
	username := c.Ctx.Input.Query("username")
	tokenString := ""
	if username != "" {
		easyToken := helpers.EasyToken{
			Username: username,
			Expires:  time.Now().Unix() + 3600,
		}
		tokenString, _ = easyToken.GenerateToken()
	}

	c.Data["json"] = tokenString
	c.ServeJSON()
}

func (c *AuthController) VerifyToken() {
	tokenString := c.Ctx.Input.Query("token")
	// tokenString := c.Ctx.Request.Header.Get("X-JWTtoken")
	easyToken := helpers.EasyToken{}
	iss, err := easyToken.ParseToken(tokenString)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
	}

	c.Data["json"] = iss
	c.ServeJSON()
}
