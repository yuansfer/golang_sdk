package controllers

import (
	"github.com/astaxie/beego"
)

type controller struct {
	beego.Controller
}

func (c *controller) checkData(key string, value interface{}) {
	if c.Data[key] = value; c.Data[key] == "" {
		c.Data[key] = "-"
	}
}
