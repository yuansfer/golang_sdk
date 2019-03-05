package controllers

import (
	"github.com/astaxie/beego"
)

type Controller struct {
	beego.Controller
}

func (c *Controller) checkData(key string, value interface{}) {
	if c.Data[key] = value; c.Data[key] == "" {
		c.Data[key] = "-"
	}
}
