package main

import (
	_ "bbs/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/videos", "video")
	beego.SetStaticPath("/books", "books")
	beego.SetStaticPath("/file", "file")
	beego.Run()
}
