package spider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

const (
	//          1  2  3  4  5  6  7  8  9  10  11
	fmtStr   = "%s|%s|  |%s|  |  |   |%s|  |%s|%s"
	fmtStr1  = "%s|%s|  |%s|  |  |[sound:%s.mp3]|%s|  |%s|%s\n" //mp3
	saveFile = "data.txt"
	ankiFile = "anki.txt"
)

type Card struct {
	Name  string
	Ph    string
	Meams string
	Ex    string
	Egen  string
	Egzh  string
}

var m map[string]*Card
var dbon bool = true

func getCard(word string) *Card {
	if !dbon {
		return nil
	}
	if c, ok := m[word]; ok {
		return c
	}
	return nil
}

func setCard(c *Card) {
	if c.Meams == "" {
		return
	}
	m[c.Name] = c
	Dump()
}

func Dump() {
	nosave := true
	if nosave {
		// return
	}

	data, err := json.Marshal(m)
	if nil != err {
		log.Println("dump data error")
	}
	ioutil.WriteFile(saveFile, data, 0644)
}

func exportCard(c *Card) string {
	line := fmt.Sprintf(fmtStr, c.Name, c.Ph, c.Meams, c.Ex, c.Egzh, c.Egen)
	return line
}

// mp3
func ExportCard(c *Card) string {
	line := fmt.Sprintf(fmtStr1, c.Name, c.Ph, c.Meams, c.Name, c.Ex, c.Egzh, c.Egen)
	return line
}

func export2Anki() {
	var lines = make([]string, 0, len(m))
	for _, card := range m {
		line := exportCard(card)
		lines = append(lines, line)
	}

	fmt.Println(m, len(m))
	write(lines)
	Dump()
}

func write(lines []string) {
	if len(lines) == 0 {
		return
	}
	data := strings.Join(lines, "\n")
	ioutil.WriteFile(ankiFile, []byte(data), 0644)
}

func saveCard(c *Card) {
	line := exportCard(c) + "\n"
	ioutil.WriteFile("test.txt", []byte(line), 0644)
}

// batch trans
func TransWords() {
	data, err := ioutil.ReadFile("input.txt")
	if nil != err {
		panic(err)
	}

	words := strings.Split(string(data), "\n")
	for _, word := range words {
		if getCard(word) != nil {
			continue
		}

		// word = word[:len(word)-1]
		p := newParser(word)
		time.Sleep(time.Second / 2)
		p.basic(word)
		p.extend(word)
		if p.c.Name == "" {
			fmt.Println("null", word, []byte(word))
		}
		m[word] = p.c
	}

	export2Anki()
}

func initDb() {
	m = make(map[string]*Card)

	data, err := ioutil.ReadFile(saveFile)
	if nil != err {
		log.Println("no data", err)
		return
	}

	err = json.Unmarshal(data, &m)
	if nil != err {
		log.Println("unmarshal json", err)
	}
}

func GetM() map[string]*Card {
	return m
}
