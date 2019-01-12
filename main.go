package main

import (
	_ "github.com/go-sql-driver/mysql"

	_ "bbs/models"
	_ "bbs/routers"

	"github.com/astaxie/beego"
)

func init() {
	beego.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	beego.Run()
}
