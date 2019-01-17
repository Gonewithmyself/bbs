package controllers

import (
	"bbs/spider"

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
	data := c.GetString("ctx")

	res := spider.Trans(data)
	if "" == res {
		c.rspFailed("no such word.")
	} else {
		c.rspSuccess(res)
	}

	// fmt.Println(data, "ctx", res)
}

func (c *MainController) rspSuccess(msg string) {
	resp := map[string]interface{}{"status": 0, "msg": "", "data": msg}
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *MainController) rspFailed(msg string) {
	resp := map[string]interface{}{"status": -1, "msg": msg}
	c.Data["json"] = resp
	c.ServeJSON()
}
