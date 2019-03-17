package main

import (
	_ "bbs/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/videos", "video")
	beego.SetStaticPath("/file", "file")
	beego.Run()
}
