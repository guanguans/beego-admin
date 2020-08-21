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
