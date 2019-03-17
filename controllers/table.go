package controllers

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"

	"github.com/astaxie/beego"
)

type TableController struct {
	beego.Controller
}

func (c *TableController) Get() {
	// c.Data["movie"] = vidioes
	c.TplName = "file.html"
}

type response struct {
	Code  int         `json:"code"`
	Count int         `json:"count"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

type item struct {
	ID    int32 `json:"id"`
	Score int32 `json:"score"`
}

func (c *TableController) Post() {
	f, h, err := c.GetFile("csv")
	if nil == err {
		fmt.Println("x", f, h.Filename, err)

	} else {
		fmt.Println("x", f, h, err)
	}

	csvf := csv.NewReader(f)
	ss, err := csvf.ReadAll()
	if nil != err {
		fmt.Println("csv", err)
	}

	fmt.Println(ss, err)
	d, err := ioutil.ReadAll(f)

	fmt.Println(c.Ctx.Request.ContentLength, err, string(d))
	c.TplName = "file.html"
}
