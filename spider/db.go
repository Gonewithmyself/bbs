package spider

import (
	"fmt"
	"io/ioutil"
)

const (
	//         1  2  3  4  5  6  7  8  9  10  11
	fmtStr = "%s|%s|  |%s|  |  |   |%s|  |%s|%s\n"
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

func flash() {

}

func getCard(word string) *Card {
	if c, ok := m[word]; ok {
		return c
	}
	return nil
}

func setCard(c *Card) {
	m[c.Name] = c
}

func saveCard(c *Card) {
	line := fmt.Sprintf(fmtStr, c.Name, c.Ph, c.Meams, c.Ex, c.Egzh, c.Egen)
	// fmt.Println(line)

	write(line)
}

func write(data string) {
	ioutil.WriteFile("test.txt", []byte(data), 0644)
}
