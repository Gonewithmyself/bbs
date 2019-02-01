package spider

import (
	"fmt"
	"regexp"
	"testing"
)

func Test_trans(t *testing.T) {

	tests := []struct {
		name string
		args string
		want string
	}{
		{"xx", "recite", "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trans(tt.args); got != tt.want {
				t.Errorf("trans() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extend(t *testing.T) {
	tests := []struct {
		name string
		args string
		want [][2]string
	}{
		{"xx", "recite", [][2]string{}},
	}
	parser := newParser("word")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if parser.extend(tt.args); len(tt.args) != -1 {
				t.Errorf("extend() = %v, want %v", "", tt.want)
			}
		})
	}
}

func Test_gjson(t *testing.T) {
	src := `{"st":[[[["wo"],["shi"]],[["ni"],["ge"]],"xinhua"], [[["ni"],["shi"]],[["wo"],["di"]],"xinhua"]]}`
	// fmt.Println(get(src, "st"))
	fmt.Println(get(src, "st.#"))

	ar := get(src, "st").Array()
	fmt.Println(len(ar))
	t.Error("")
}

func BenchmarkTrans(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Trans("silence")
	}
}

func Test_getEx(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{`{"word_third":["gets"],"word_done":["got"," gotten"],"word_pl":"","word_est":"","word_ing":["getting"],"word_er":"","word_past":["got"]}`}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getEx(tt.args.src); got != tt.want {
				t.Errorf("getEx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getEx(src string) string {
	exPatt = regexp.MustCompile(`\[([^/[]+?)\]`)
	src = omitQuote(src)
	fmt.Println(exPatt.FindAllString(src, -1))
	return exPatt.String()
}

func Test_to(t *testing.T) {
	TransWords()
	t.Error("123")
}

func Test_mp3(t *testing.T) {
	getMp3("tier")
	t.Error("123")
}

func Test_sign(t *testing.T) {
	getSign("over")
	t.Error("123")
}

func Test_token(t *testing.T) {
	prepareToken()
	t.Log(token, gtk)
	t.Error("123")
}

func Test_regex(t *testing.T) {
	var src = `goto school, goto home, goto station`
	patt := regexp.MustCompile(`goto ([\S]+?)`)
	res := patt.FindAllStringSubmatch(src, -1)
	t.Log(res)

	t.Error("123")
}
