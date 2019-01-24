package spider

import (
	"fmt"
	"log"
	"runtime/debug"
	"strings"

	"github.com/tidwall/gjson"
)

// external
func Trans(word string) string {
	c := getCard(word)
	if nil == c {
		c = trans(word)
	}

	if c.Name == "" {
		return ""
	}
	// for test
	// saveCard(c)

	res := fmt.Sprintf("%s%s%s", c.Ph, Ln, escapeHtml(c.Meams))
	return res
}

func trans(word string) *Card {
	parser := newParser(word)
	parser.basic(word)
	if "" == parser.c.Name {
		return parser.c
	}

	go func() {
		parser.extend(word)
		getMp3(word)
		setCard(parser.c)
	}()

	return parser.c
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
			log.Printf("stack %s\n", debug.Stack())
			sp.c.Name = ""
		}
	}()

	str := gjson.Get(body, "dict").String()
	if str == "[]" {
		sp.c.Name = ""
		log.Printf("parse dict [] %s\nbody %s\n", str, body)
	}

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
	items := get(body, "data.edict.item.#.tr_group").Array()
	groups := toString(items)

	for i := range groups {
		// just extract one eg per group
		tr := get(groups[i], "0.tr").String()
		eg := get(groups[i], "0.example").String()
		line := fmt.Sprintf("%d. %s<br/>   eg.: %s<br/>", i+1, tr, eg)
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
		sp.c.Egzh += fmt.Sprintf("%d. %s<br/>   %s<br/>", i+1, en, zh)
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
		res = fmt.Sprintf("%d. %-8s %-2s<br/>", i+1, attr, res)
		sp.c.Meams += res
	}

	sp.c.Meams = replHtmlSpace(sp.c.Meams)
}

// ex : {"word_third":["tests"],"word_done":["tested"],"word_pl":["tests"],"word_est":""}
func (sp *parser) parseEx(d string) {
	ex := gjson.Get(d, "exchange").String()

	parts := exPatt.FindAllString(ex, -1)
	sp.c.Ex = omitQuote(strings.Join(parts, " "))
	// fmt.Println("ex", parts)
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
