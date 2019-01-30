package spider

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/parnurzeal/gorequest"
	"github.com/robertkrimen/otto"
	"github.com/tidwall/gjson"
)

const (
	url0        = "https://fanyi.baidu.com/basetrans"
	url1        = "https://fanyi.baidu.com/extendtrans"
	url2        = "http://localhost:8080"
	Url         = url1
	Agent       = "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Mobile Safari/537.36"
	JsonFmt     = `from=en&to=zh&query=%s&token=05301f0daa555e723f477ec3b63f3638&sign=%s`
	phFmt       = `https://fanyi.baidu.com/gettts?lan=en&text=%s&spd=3&source=wise`
	Ln          = "<br/>"
	exPattSting = `\[([\s\S]+?)\]`
	mediaDir    = `./media/%s.mp3`
	cookie      = `BAIDUID=26573F942C2978C7C6FAA992BD1C75C8:FG=1;locale=zh;path=/;domain=.baidu.com`
	SignJs      = "spider/sign.js"
)

type C struct {
	From  string
	To    string
	Query string
}

var agent *gorequest.SuperAgent
var phAgent *gorequest.SuperAgent
var agents [2]*gorequest.SuperAgent
var exPatt *regexp.Regexp

func init() {
	for i := range agents {
		url := url0
		if i == 1 {
			url = url1
		}
		agents[i] = gorequest.New().Post(url).Set("User-Agent", Agent).Set("X-Requested-With", "XMLHttpRequest").Set("Connection", "keep-alive").
			Set("Cookie", cookie).Set("Referer", "https://fanyi.baidu.com/")
	}

	phAgent = gorequest.New().Set("User-Agent", Agent).Set("Connection", "keep-alive").
		Set("Referer", "https://fanyi.baidu.com/").
		Set("Cookie", "BAIDUID=26573F942C2978C7C6FAA992BD1C75C8:FG=1;")

	exPatt = regexp.MustCompile(exPattSting)
	initDb()
}

func post(word, url string) string {
	data := fmt.Sprintf(JsonFmt, word, getSign(word))
	if "b" == url {
		agent = agents[0]
	} else {
		agent = agents[1]
	}

	// setQuery(word)
	resp, body, err := agent.SendString(data).End()
	//	fmt.Println("data:", agent.Data)

	if nil != err {
		log.Print(err)
		return ""
	}

	_ = resp
	return body
}

func getMp3(word string) {
	url := fmt.Sprintf(phFmt, word)
	_, data, err := phAgent.Get(url).EndBytes()

	if nil != err {
		log.Println("get mp3", err)
		return
	}

	file := fmt.Sprintf(mediaDir, word)
	if err := ioutil.WriteFile(file, data, 0755); nil != err {
		log.Println(err)
	}
}

func GetMp3(word string) {
	getMp3(word)
}

func setQuery(word string) {
	agent.Data["query"] = word
	agent.Data["from"] = "en"
	agent.Data["to"] = "zh"
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

func replHtmlSpace(src string) string {
	return strings.Replace(src, " ", "&nbsp;", -1)
}

func omitQuote(src string) string {
	return strings.Replace(src, "\"", "", -1)
}

func recoverFn(params ...interface{}) {
	r := recover()
	if nil != r {
		log.Println(params, r)
	}
}

func getSign(word string) string {
	jsWord, err := vm.ToValue(word)
	if nil != err {
		panic(err)
	}

	res, err := vm.Call("token", nil, jsWord, gtk)
	if nil != err {
		panic(err)
	}

	sign, _ := res.ToString()
	return sign
}

// init vm
var vm *otto.Otto
var gtk otto.Value

func init() {
	vm = otto.New()
	src, err := ioutil.ReadFile(SignJs)
	if nil != err {
		panic(err)
	}

	_, err = vm.Run(src)
	if nil != err {
		panic(err)
	}

	gtk, err = vm.ToValue("320305.131321201")
	if nil != err {
		panic(err)
	}
}
