package AdminControllers

import (
	"time"

	"beego-admin/helpers"
)

type AuthController struct {
	BaseController
}

func (c *AuthController) GetToken() {
	username := c.Ctx.Input.Query("username")
	tokenString := ""
	if username != "" {
		et := helpers.EasyToken{
			Username: username,
			Expires:  time.Now().Unix() + 3600, //Segundos
		}
		tokenString, _ = et.GenerateToken()
	}

	c.Data["json"] = tokenString
	c.ServeJSON()
}

func (c *AuthController) VerifyToken() {
	tokenString := c.Ctx.Input.Query("token")
	// O puede ser leído de una cabecera HEADER!!
	// tokenString := c.Ctx.Request.Header.Get("X-JWTtoken")
	et := helpers.EasyToken{}
	valido, _, _ := et.ValidateToken(tokenString)
	if !valido {
		c.Ctx.Output.SetStatus(401)
		c.Data["json"] = "token验证失败"
		c.ServeJSON()
	}
	c.Data["json"] = "token验证成功"
	c.ServeJSON()
}
