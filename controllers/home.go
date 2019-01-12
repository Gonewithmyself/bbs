package controllers

import "fmt"

type HomeController struct {
	baseController
}

func (c *HomeController) Home() {
	c.Data["login"] = c.sess.DoesLogin()
	fmt.Println(c.sess.DoesLogin(), c.sess.Get())
	c.TplName = "index.html"
}
