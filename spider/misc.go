package spider

import (
	"fmt"
	"log"
	"strings"

	"github.com/parnurzeal/gorequest"
	"github.com/tidwall/gjson"
)

const (
	url1    = "https://fanyi.baidu.com/basetrans"
	url0    = "https://fanyi.baidu.com/extendtrans"
	url2    = "http://localhost:8080"
	Url     = url1
	Agent   = "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Mobile Safari/537.36"
	JsonFmt = `from=en&to=zh&query=%s`
	Ln      = "<br />"
)

type C struct {
	From  string
	To    string
	Query string
}

type Ext struct {
	St    [][2]string
	Edict map[string]string
}

func NewExt() *Ext {
	return &Ext{
		Edict: make(map[string]string),
	}
}

var req = gorequest.New()

func post(word, url string) string {
	if "b" == url {
		url = url1
	} else {
		url = url0
	}

	data := fmt.Sprintf(JsonFmt, word)
	resp, body, err := req.Post(url).Set("User-Agent", Agent).Set("X-Requested-With", "XMLHttpRequest").
		Set("Cookie", "BAIDUID=3395EB2E85B1B8F49DF5F52818DFEFE4:FG=1;").Set("Referer", "https://fanyi.baidu.com/").
		Send(data).End()

	if nil != err {
		log.Print(err)
		return ""
	}

	_ = resp
	return body
}

func get(js, path string) gjson.Result {
	return gjson.Get(js, path)
}

func toString(list []gjson.Result) (res []string) {
	for i := range list {
		res = append(res, list[i].String())
	}
	return res
}

var jsEMap = map[string]string{
	"\\u003c": "<",
	"\\u003e": ">",
	"\\u0026": "&",
}

func escapeHtml(content string) string {
	for k, v := range jsEMap {
		content = strings.Replace(content, k, v, -1)
	}
	return content
}

func recoverFn(params ...interface{}) {
	r := recover()
	if nil != r {
		log.Println(params, r)
	}
}
