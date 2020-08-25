package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

type ErrorJsonData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type SuccessJsonData struct {
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
	Status string      `json:"status"`
}

func (c *BaseController) SuccessJson(data interface{}) {
	c.Data["json"] = SuccessJsonData{200, data, "success"}
	c.ServeJSON()
}

func (c *BaseController) ErrorJson(message string) {
	c.Data["json"] = ErrorJsonData{422, message, "error"}
	c.ServeJSON()
}
