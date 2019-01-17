package controllers

import (
	"bbs/lib/types"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type baseController struct {
	beego.Controller
	o              orm.Ormer
	sess           *types.Session
	fn             func()
	controllerName string
	actionName     string
	auth           string
}

func (p *baseController) Prepare() {
	controllerName, actionName := p.GetControllerAndAction()
	p.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	p.actionName = strings.ToLower(actionName)
	p.o = orm.NewOrm()

	p.initSession()
	p.SetLogin()

}

func (c *baseController) SetLogin() {
	c.Data["login"] = c.sess.DoesLogin()
}

func (p *baseController) History(msg string, url string) {
	if url == "" {
		p.Ctx.WriteString("<script>alert('" + msg + "');window.history.go(-1);</script>")
		p.StopRun()
	} else {
		p.Redirect(url, 302)
	}
}

//获取用户IP地址
func (p *baseController) getClientIp() string {
	s := strings.Split(p.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

//获取用户IP地址
func (c *baseController) initSession() {
	var (
		sess *types.Session
		ok   bool
		auth = c.Ctx.GetCookie("_s")
	)

	if auth != "" {
		isess := c.GetSession(auth)
		if sess, ok = isess.(*types.Session); !ok || nil == sess {
			sess = &types.Session{}
		}

	} else {
		auth = c.genCookie()
		c.Ctx.SetCookie("_s", auth, 600)
		sess = &types.Session{}
	}

	c.sess = sess
	c.auth = auth
	c.SetSession(auth, sess)
}

func (c *baseController) delSession() {
	c.SetSession(c.auth, nil)
}

//获取用户IP地址
func (c *baseController) genCookie() string {
	ts := time.Now().UnixNano()
	return c.getClientIp() + strconv.FormatInt(ts, 10)
}
