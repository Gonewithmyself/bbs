package agent

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

var (
	lsp = &lAgent{}
	bp  = NewBallParser()
)

func init() {
	lsp.SetReuse(true)
	lsp.SetHeader(LoadHeader("conf.ini"))
	lsp.SetCallback(lsp.Do)
	loadLottery()
}

func GetLotteryInfo() {
	// http://kaijiang.500.com/shtml/ssq/19016.shtml
	lsp.SetCallback(lsp.getPages)
	lsp.Get("http://kaijiang.500.com/shtml/ssq/19016.shtml")

	lsp.SetCallback(lsp.Do)
	lsp.Parse()
}

type ltInfo struct {
	ball string
	l    int
}

type ltslice []*ltInfo

func (l ltslice) Len() int           { return len(l) }
func (l ltslice) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ltslice) Less(i, j int) bool { return l[i].l < l[j].l }

func AnalyseLottery() {
	var (
		m = make([]*ltInfo, len(bp.Data))
		i int
	)
	for k, v := range bp.Data {
		m[i] = &ltInfo{k, len(v)}
		i++
	}

	sort.Sort(ltslice(m))
	for idx, item := range m {
		fmt.Println(item)
		if idx > 100 {
			break
		}
	}
}

type lAgent struct {
	Agent
	pages []string
}

func (self *lAgent) Do() {
	defer func() {
		r := recover()
		if nil != r {
			fmt.Println(self.errs)
			fmt.Println(self.resp)
			panic(r)
		}
	}()
	r, err := gzip.NewReader(self.resp.Body)
	body, err := ioutil.ReadAll(r)
	panicError(err)
	src := string(body)
	bp.Parse(src)
}

func (self *lAgent) Parse() {
	fmt.Println(len(bp.Data), len(bp.PhaseInfo), len(self.pages))
	for idx, page := range self.pages {
		if bp.DoesParsed(page) {
			continue
		}

		url := fmt.Sprintf("http://kaijiang.500.com/shtml/ssq/%s.shtml", page)
		self.Get(url)

		if 0 != idx && idx%5 == 0 {
			go dumpLottery()
		}
	}

	// fmt.Println(bp, "self.pages")
	dumpLottery()
}

func (lsp *lAgent) getPages() {
	r, err := gzip.NewReader(lsp.resp.Body)
	body, err := ioutil.ReadAll(r)
	panicError(err)
	src := string(body)
	// fmt.Println(string("body"), "error\n", err)

	lsp.pages = bp.ParsePages(src)
}

type BallParser struct {
	red, blue, phase, pages *regexp.Regexp
	Data                    map[string][]string
	PhaseInfo               map[string]string
}

func NewBallParser() *BallParser {
	return &BallParser{
		red:       regexp.MustCompile(`<li class="ball_red">([0-9]+)</li>`),
		blue:      regexp.MustCompile(`<li class="ball_blue">([0-9]+)</li>`),
		phase:     regexp.MustCompile(`<font class="cfont2"><strong>([0-9]+)</strong></font>`),
		pages:     regexp.MustCompile(`/shtml/ssq/([0-9]+).shtml`),
		Data:      make(map[string][]string),
		PhaseInfo: make(map[string]string),
	}
}

func (p *BallParser) ParsePages(src string) []string {
	group := p.pages.FindAllStringSubmatch(src, -1)
	return extractGroup(group)
}

func (p *BallParser) DoesParsed(page string) bool {
	_, ok := p.PhaseInfo[page]
	return ok
}

func (p *BallParser) Parse(src string) {
	k, v := p.parsePhase(src), p.parseBall(src)
	if _, ok := p.PhaseInfo[k]; !ok {
		p.PhaseInfo[k] = v
	} else {
		return
	}

	phases := p.Data[v]
	for _, e := range phases {
		if e == k {
			return
		}
	}
	phases = append(phases, k)
	p.Data[v] = phases
}

func (p *BallParser) parsePhase(src string) string {
	group := p.phase.FindAllStringSubmatch(src, -1)
	return strings.Join(extractGroup(group), "")
}

func (p *BallParser) parseBall(src string) string {
	group := p.red.FindAllStringSubmatch(src, -1)
	res := extractGroup(group)

	group = p.blue.FindAllStringSubmatch(src, -1)
	res = append(res, extractGroup(group)...)
	return strings.Join(res, " ")
}

func extractGroup(src [][]string) []string {
	var res = make([]string, len(src))
	for idx, one := range src {
		res[idx] = one[1]
	}
	return res
}

func dumpLottery() {
	data, err := json.Marshal(bp)
	panicError(err)
	ioutil.WriteFile("lottery.json", data, 0644)
}

func loadLottery() {
	data, err := ioutil.ReadFile("lottery.json")
	panicError(err)
	err = json.Unmarshal(data, bp)
	panicError(err)
}
