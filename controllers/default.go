package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//var test models.Test
	//test.Data = rand.Int()
	//c.Data["id"] = test.Id
	//c.Data["data"] = test.Data
	name := c.GetString("name")
	c.Data["Website"] = name
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

