package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}

func (c *MainController) Post() {
	// c.TplName = "index.html"
	data := c.GetString("from")

	// d := spider.C{}
	// c.ParseForm(d)
	// data := d.From
	// fmt.Println(d)

	fmt.Println(data, "ctx")
	resp := map[string]interface{}{"status": 0, "msg": "", "data": data}
	c.Data["json"] = resp
	c.ServeJSON()
}
