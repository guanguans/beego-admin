// @APIVersion 1.0.0
// @Title beego-admin Demo API
// @Contact ityaozm@gmail.com
// @TermsOfServiceUrl https://github.com/guanguans
// @License MIT
// @LicenseUrl https://github.com/guanguans/beego-admin/blob/master/LICENSE
package routers

import (
	"beego-admin/controllers/api"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("api/v1",
		beego.NSRouter("demo", &api.DemoController{}, "*:Demo"),
	)
	beego.AddNamespace(ns)
}
