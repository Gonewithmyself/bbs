package controllers

import (
	"os"
	"path/filepath"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["movie"] = vidioes
	c.TplName = "list.html"
}

func (c *MainController) Play() {
	c.Data["name"] = c.GetString("name")
	c.TplName = "video.html"
}

var vidioes []string

func init() {
	listVidios()
}

func listVidios() []string {
	//filepath.WalkFunc
	filepath.Walk("video", func(name string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		vidioes = append(vidioes, info.Name())
		return nil
	})
	return nil
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
