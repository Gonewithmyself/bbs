package spider

import (
	"encoding/json"
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
)

type C struct {
	From  string
	To    string
	Query string
}

type Dict struct {
	Ph    string
	Ex    string
	Means map[string]string
}

type Ext struct {
	St    [][2]string
	Edict map[string]string
}

func NewDict() *Dict {
	return &Dict{
		Means: make(map[string]string),
	}
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

func basic(word string) *Dict {
	body := post(word, "b")
	var dict *Dict
	defer func() {
		r := recover()
		if nil != r {
			log.Print(r)
			dict = nil
		}
	}()

	dict = NewDict()
	str := gjson.Get(body, "dict").String()
	dict.Ex = parseEx(str)
	dict.Means = parseMeans(str)
	dict.Ph = get(str, "symbols.0.ph_am").String()

	return dict
}

func extend(word string) *Ext {
	body := post(word, "")
	var ext *Ext
	defer func() {
		r := recover()
		if nil != r {
			log.Print(r)
			ext = nil
		}
	}()
	ext = NewExt()
	ext.Edict = extendEn(body)
	ext.St = extendZh(body)
	return ext
}

// en eg.
func extendEn(body string) map[string]string {
	items := get(body, "data.edict.item.0.tr_group").Array()
	groups := toString(items)

	var edict = make(map[string]string, len(groups))
	for i := range groups {
		tr := get(groups[i], "tr").String()
		eg := get(groups[i], "example").String()
		edict[tr] = eg
	}
	return edict
}

// zh eg.
func extendZh(body string) [][2]string {
	st := get(body, "data.st").String()
	books := get(st, "#.2").Array()
	ens := get(st, "#.0").Array()
	zhs := get(st, "#.1").Array()

	res := make([][2]string, len(books))
	idx := 0
	for i := range books {
		if !strings.Contains(books[i].String(), "《") {
			continue
		}

		res[i][0] = parseEg(ens[i].String(), true)
		res[i][1] = parseEg(zhs[i].String(), false)
		idx++
	}

	return res[:idx]
}

func parseEg(s string, en bool) string {
	ss := get(s, "#.0").Array()
	words := toString(ss)

	tag := ""
	if en {
		tag = " "
	}

	line := strings.Join(words, tag)
	return line
}

func Trans(word string) string {
	dict := basic(word)
	if nil == dict {
		return ""
	}

	d, _ := json.Marshal(dict.Means)
	res := escapeHtml(string(d)) + "\r\n" + dict.Ph
	go flash(word, dict)
	return res
}

func parseMeans(d string) map[string]string {
	parts := gjson.Get(d, "symbols.0.parts").Array()
	m := make(map[string]string)
	for _, part := range parts {
		p := part.String()
		attr := get(p, "part").String()
		means := get(p, "means").Array()
		mm := toString(means)

		res := strings.Join(mm, ",")
		m[attr] = res
	}

	return m
}

// ex : {"word_third":["tests"],"word_done":["tested"],"word_pl":["tests"],"word_est":""}
func parseEx(d string) string {
	ex := gjson.Get(d, "exchange").String()

	parts := strings.Split(ex, ",")
	res := make([]string, 0, 5)
	for _, part := range parts {
		p := strings.Split(part, ":")[1]
		if p == `""` {
			continue
		}

		p = strings.TrimLeft(p, "[\"")
		p = strings.TrimRight(p, "}")
		p = strings.TrimRight(p, ",\"]")
		res = append(res, p)
	}

	return strings.Join(res, ",")
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
