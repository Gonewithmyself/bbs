package main

import (
	_ "bbs/routers"
	"bbs/spider"
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {

	t.Parallel()
	spider.TransWords()
	t.Error("123")
}

func Test_anki(t *testing.T) {
	// data, err := ioutil.ReadFile("anki.txt")
	// t.Error("read", err)

	// var words []string
	// lines := strings.Split(string(data), "\n")
	// for _, line := range lines {
	// 	word := strings.Split(line, "|")[0]
	// 	// t.Log(word)
	// 	words = append(words, word)
	// 	_ = word
	// }

	// m := spider.GetM()
	// for k := range m {
	// 	//t.Log("xx", k)
	// 	words = append(words, k)
	// }

	// dd := strings.Join(words, "\n")
	// ioutil.WriteFile("input.txt", []byte(dd), 644)

	spider.TransWords()
	t.Error("xx")
}

func Test_mp3(t *testing.T) {
	m := spider.GetM()
	for k := range m {
		if _, ok := mbak[k]; ok {
			continue
		}
		spider.GetMp3(k)
	}

	t.Error("xx")
}

func Test_sptmp3(t *testing.T) {
	m := spider.GetM()
	lines := []string{}
	for word, card := range m {
		if _, ok := mbak[word]; ok {
			continue
		}

		line := spider.ExportCard(card)
		lines = append(lines, line)
	}

	data := strings.Join(lines, "")
	ioutil.WriteFile("anki.txt", []byte(data), 0644)
	t.Error("xx", len(mbak), len(m))
}

var mbak map[string]*spider.Card

func init() {
	mbak = make(map[string]*spider.Card)

	data, err := ioutil.ReadFile("data_bak.txt")
	if nil != err {
		panic(err)
	}

	err = json.Unmarshal(data, &mbak)
	if nil != err {
		panic(err)
	}
}
