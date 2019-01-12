package controllers

type HomeController struct {
	baseController
}

func (c *HomeController) Home() {
	// c.SetLogin()
	c.TplName = "index.html"
}
