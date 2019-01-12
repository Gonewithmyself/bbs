package controllers

import (
	"bbs/lib/types"
	"bbs/util"
	"fmt"
	"time"
)

type UserController struct {
	baseController
}

//配置信息
func (c *UserController) Login() {
	switch c.Ctx.Request.Method {
	case "GET":
		c.loginGet()
		println("get")
	case "POST":
		c.loginPost()
		println("post")
	default:
		c.History("", "/")
		println("defualt")
	}

	println("ctl name", c.controllerName)

}

//配置信息
func (c *UserController) Reg() {
	switch c.Ctx.Request.Method {
	case "GET":
		c.regGet()
		println("get")
	case "POST":
		c.regPost()
		println("post")
	default:
		c.History("", "/")
		println("defualt")
	}

}

func (c *UserController) regGet() {
	c.TplName = c.controllerName + "/reg.html"
}

func (c *UserController) regPost() {
	data := &types.RegForm{}
	err := c.ParseForm(data)
	if nil != err || data.Account == "" {
		str := data.Account
		if nil != err {
			str = err.Error()
		}
		c.msgString(-1, str)
		return
	}

	fmt.Println(data)
	if data.Pass == "" || data.Pass != data.Repass {
		c.msgString(-1, "密码不相同")
		return
	}

	user := c.getUser()
	if nil != user {
		c.msgString(-1, "短时间内无法多次注册")
		return
	}
	user = &types.User{}

	user.Account = data.Account
	err = c.o.Read(user, "account")
	if nil == err {
		c.msgString(-1, "already regist")
		return
	}

	user.Password = util.Md5(data.Pass)
	user.Name = data.Name
	user.Created = time.Now()

	_, err = c.o.Insert(user)
	if nil != err {
		c.msgString(-1, err.Error())
		return
	}

	c.msgString(0, "login")
}

func (c *UserController) msgString(code int, msg interface{}) {
	var s string
	if v, ok := msg.(string); ok {
		s = v
	} else if v, ok := msg.(error); ok {
		s = v.Error()
	} else {
		fmt.Println(msg)
		s = "unknown error"
	}
	c.Data["json"] = map[string]interface{}{"status": code, "msg": s}
	c.ServeJSON()
}

func (c *UserController) loginGet() {
	c.TplName = c.controllerName + "/login.html"
	fmt.Println(c.TplName)
}

func (c *UserController) loginPost() {
	user := c.getUser()
	if nil != user {
		c.msgString(0, "/")
		return
	}

	user = &types.User{}
	data := &types.RegForm{}
	err := c.ParseForm(data)
	if nil != err {
		c.msgString(-1, err)
		return
	}
	user.Account = data.Account

	err = c.o.Read(user, "account")
	if nil != err {
		c.msgString(-1, err)
		return
	}

	if util.Md5(data.Pass) != user.Password {
		c.msgString(-1, "wrong password")
		fmt.Println(data, user, util.Md5(data.Pass))
		return
	}

	c.sess.Set(user)
	c.sess.SetAuth(true)
	c.msgString(0, "/")
}

func (c *UserController) getUser() *types.User {
	return c.sess.User
}
