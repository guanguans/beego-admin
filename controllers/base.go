package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

type JsonData struct {
	Code int
	Data interface{}
	Msg  string
}

func (c *BaseController) AjaxSuccess(data interface{}) {
	jsonData := JsonData{1, data, "ok"}
	c.AjaxReturn(jsonData)
}

func (c *BaseController) AjaxError(msg string) {
	jsonData := JsonData{-1, nil, msg}
	c.AjaxReturn(jsonData)
}

func (c *BaseController) AjaxReturn(jsonData JsonData) {
	c.Data["json"] = jsonData
	c.ServeJSON()
}
