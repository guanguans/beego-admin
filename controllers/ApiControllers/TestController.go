package ApiControllers

import "beego-admin/controllers"

type BaseController struct {
	controllers.BaseController
}

type BeegoAdmin struct {
	ProjectName string
	ProjectUrl string
	Author string
}

func (c *BaseController) Test()  {
	BeegoAdmin := BeegoAdmin{"beego-admin", "https://github.com/guanguans/beego-admin","guanguans"}
	c.Data["json"] = BeegoAdmin
	c.ServeJSON()
}