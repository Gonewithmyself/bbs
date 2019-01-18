package spider

type Card struct {
	Name  string
	Ph    string
	Meams []string
	Ex    string
	Egen  []string
	Egzh  []string
}

var m map[string]*Card

func flash(word string, d *Dict) {
	ext := extend(word)

	_ = ext
}
