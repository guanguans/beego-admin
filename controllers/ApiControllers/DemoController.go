package ApiControllers

type DemoController struct {
	BaseController
}

type BeegoAdmin struct {
	ProjectName string
	ProjectUrl string
	Author string
}

func (this *BaseController) Demo()  {
	BeegoAdmin := BeegoAdmin{"beego-admin", "https://github.com/guanguans/beego-admin","guanguans"}
	this.Data["json"] = BeegoAdmin
	this.ServeJSON()
}