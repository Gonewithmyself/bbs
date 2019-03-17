package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
)

type BookController struct {
	beego.Controller
}

func (b *BookController) Get() {
	list := listFiles("books")
	if id := b.GetString("id"); "" == id {
		b.Data["list"] = list
		b.TplName = "books.html"
	} else {
		i, _ := strconv.Atoi(id)
		b.Data["Name"] = list[i]
		b.TplName = "book.html"
	}
}
