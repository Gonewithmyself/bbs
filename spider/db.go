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
	//         1  2  3  4  5  6  7  8  9  10  11
	fmtStr   = "%s|%s|  |%s|  |  |   |%s|  |%s|%s"
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
var dbon bool

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
	m[c.Name] = c
	Dump()
}

func Dump() {
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

func export2Anki() {
	var lines = make([]string, 0, len(m))
	for _, card := range m {
		line := exportCard(card)
		lines = append(lines, line)
	}

	fmt.Println(m, len(m))
	write(lines)
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
		p := newParser(word)
		time.Sleep(time.Second)
		p.basic(word)
		p.extend(word)
		if p.c.Name == "" {
			fmt.Println("null", word)
		}
		m[word] = p.c
		fmt.Println("now", word)
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
