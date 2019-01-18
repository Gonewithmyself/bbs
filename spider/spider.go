package spider

import (
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

	res := fmt.Sprintf("%s%s%s", parser.c.Ph, Ln, escapeHtml(parser.c.Meams))
	go func() {
		parser.extend(word)
		saveCard(parser.c)
	}()
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
	sp.c.Ph = fmt.Sprintf("[%s]", get(str, "symbols.0.ph_am").String())
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

	for i := range groups {
		tr := get(groups[i], "tr").String()
		eg := get(groups[i], "example").String()
		line := fmt.Sprintf("%d. %s<br />   eg.: %s<br />", i+1, tr, eg)
		sp.c.Egen += line
	}
}

// zh eg.
func (sp *parser) extendZh(body string) {
	st := get(body, "data.st").String()
	books := get(st, "#.2").Array()
	ens := get(st, "#.0").Array()
	zhs := get(st, "#.1").Array()

	for i := range books {
		if !strings.Contains(books[i].String(), "《") {
			break
		}

		en := parseEg(ens[i].String(), true)
		zh := parseEg(zhs[i].String(), false)
		sp.c.Egzh += fmt.Sprintf("%d. %s<br />   %s<br />", i+1, en, zh)
	}
}

func (sp *parser) parseMeans(d string) {
	parts := gjson.Get(d, "symbols.0.parts").Array()
	for i, part := range parts {
		p := part.String()
		attr := get(p, "part").String()
		means := get(p, "means").Array()
		mm := toString(means)

		res := strings.Join(mm, ",")
		res = fmt.Sprintf("%d. %-5s %s<br />", i+1, attr, res)
		sp.c.Meams += res
	}
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
