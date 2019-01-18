package spider

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/tidwall/gjson"
)

// external
func Trans(word string) string {
	parser := newParser(word)
	parser.basic(word)
	if nil == parser.c {
		return ""
	}

	d, _ := json.Marshal(parser.c.Meams)
	res := escapeHtml(string(d)) + "\r\n" + parser.c.Ph
	//	go flash(word, dict)
	return res
}

type parser struct {
	c *Card
}

func newParser(word string) *parser {
	return &parser{
		c: &Card{Name: word},
	}
}

func (sp *parser) basic(word string) {
	body := post(word, "b")
	defer func() {
		r := recover()
		if nil != r {
			log.Print(r)
			sp.c = nil
		}
	}()

	str := gjson.Get(body, "dict").String()

	sp.parseEx(str)
	sp.parseMeans(str)
	sp.c.Ph = get(str, "symbols.0.ph_am").String()
}

func (sp *parser) extend(word string) {
	body := post(word, "")
	defer func() {
		r := recover()
		if nil != r {
			log.Print(r)
		}
	}()
	sp.extendEn(body)
	sp.extendZh(body)
}

// en eg.
func (sp *parser) extendEn(body string) {
	items := get(body, "data.edict.item.0.tr_group").Array()
	groups := toString(items)

	sp.c.Egen = make([]string, len(groups))
	for i := range groups {
		tr := get(groups[i], "tr").String()
		eg := get(groups[i], "example").String()
		line := fmt.Sprintf("meaning: %s, eg.: %s", tr, eg)
		sp.c.Egen[i] = line
	}
}

// zh eg.
func (sp *parser) extendZh(body string) {
	st := get(body, "data.st").String()
	books := get(st, "#.2").Array()
	ens := get(st, "#.0").Array()
	zhs := get(st, "#.1").Array()

	sp.c.Egzh = make([]string, len(books))
	idx := 0
	for i := range books {
		if !strings.Contains(books[i].String(), "《") {
			break
		}

		en := parseEg(ens[i].String(), true)
		zh := parseEg(zhs[i].String(), false)
		sp.c.Egzh[idx] = fmt.Sprintf("%s   %s", en, zh)
		idx++
	}

	sp.c.Egzh = sp.c.Egzh[:idx]
}

func (sp *parser) parseMeans(d string) map[string]string {
	parts := gjson.Get(d, "symbols.0.parts").Array()
	m := make(map[string]string)
	sp.c.Meams = make([]string, len(parts))
	for i, part := range parts {
		p := part.String()
		attr := get(p, "part").String()
		means := get(p, "means").Array()
		mm := toString(means)

		res := strings.Join(mm, ",")
		m[attr] = res
		sp.c.Meams[i] = res
	}

	return m
}

// ex : {"word_third":["tests"],"word_done":["tested"],"word_pl":["tests"],"word_est":""}
func (sp *parser) parseEx(d string) {
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

	sp.c.Ex = strings.Join(res, ",")
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
