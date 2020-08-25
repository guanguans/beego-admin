package api

type DemoController struct {
	BaseController
}

type BeegoAdmin struct {
	ProjectName string
	ProjectUrl  string
	Author      string
}

func (c *BaseController) Demo() {
	BeegoAdmin := BeegoAdmin{"beego-admin", "https://github.com/guanguans/beego-admin", "guanguans"}
	c.Data["json"] = BeegoAdmin
	c.ServeJSON()
}
