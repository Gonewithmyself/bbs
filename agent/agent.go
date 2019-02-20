package agent

import (
	"io/ioutil"
	"strings"

	"github.com/parnurzeal/gorequest"
)

type IAgent interface {
	Do()
}

type Agent struct {
	agent *gorequest.SuperAgent
	resp  gorequest.Response
	body  string
	errs  []error
	cb    func()
}

func LoadHeader(confFile string) map[string]string {
	data, err := ioutil.ReadFile(confFile)
	if nil != err {
		panic(err)
	}

	m := make(map[string]string)
	lines := strings.Split(string(data), "\r\n")
	for _, line := range lines {
		res := strings.Split(line, ":")
		m[res[0]] = res[1][1:]
	}

	return m
}

func NewAgent() *Agent {
	sp := &Agent{}
	sp.SetReuse(true)
	return sp
}

func (self *Agent) SetHeader(m map[string]string) {
	for k, v := range m {
		self.agent.Set(k, v)
	}
}

func (self *Agent) SetReuse(b bool) {
	if nil == self.agent {
		self.agent = gorequest.New()
	}
	self.agent.SetDoNotClearSuperAgent(b)
}

func (self *Agent) SetCallback(cb func()) {
	self.cb = cb
}

func (self *Agent) Post(url string) {
	self.request("post", url)
}

func (self *Agent) Get(url string) {
	self.request("get", url)
}

func (self *Agent) Head(url string) {
	self.request("head", url)
}

func (self *Agent) Delete(url string) {
	self.request("delete", url)
}

// do some magic
func (self *Agent) Do() {
}

func (self *Agent) _do(re gorequest.Response, sf string, de []error) {
	self.resp, self.body, self.errs = re, sf, de
	self.cb()

}

func (self *Agent) request(method, url string) {
	switch method {
	case "get":
		self.agent = self.agent.Get(url)
	case "post":
		self.agent = self.agent.Post(url)
	case "head":
		self.agent = self.agent.Head(url)
	case "delete":
		self.agent = self.agent.Delete(url)
	default:
		panic("unknown method")
	}

	self.agent.End(self._do)
}

func (self *Agent) requestSafe(method, url string) {
	self.request(method, url)
	self.agent.ClearSuperAgent()
}
