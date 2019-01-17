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

func NewDict() *Dict {
	return &Dict{
		Means: make(map[string]string),
	}
}

func trans(word string) *Dict {
	req := gorequest.New()
	data := fmt.Sprintf(JsonFmt, word)
	resp, body, err := req.Post(Url).Set("User-Agent", Agent).Set("X-Requested-With", "XMLHttpRequest").
		Set("Cookie", "BAIDUID=3395EB2E85B1B8F49DF5F52818DFEFE4:FG=1;").Set("Referer", "https://fanyi.baidu.com/").
		Send(data).End()

	if nil != err {
		log.Print(err)
	}

	dict := NewDict()
	str := gjson.Get(body, "dict").String()
	dict.Ex = parseEx(str)
	dict.Means = parseMeans(str)
	dict.Ph = get(str, "symbols.0.ph_am").String()

	_ = resp
	return dict
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

		println(p, "xxx")
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
