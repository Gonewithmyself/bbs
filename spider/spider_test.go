package spider

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
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
		want string
	}{
		{"xx", "recite", "nil"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extend(tt.args); got.St[0][0] != tt.want {
				t.Errorf("extend() = %v, want %v", got.St, tt.want)
			}
		})
	}
}

func Test_u2c(t *testing.T) {
	src := "\u4ed6\u4eec\u4e92\u76f8\u6717\u8bf5\u8bd7"
	t1 := "他们互相朗诵诗"
	d1 := strconv.QuoteToASCII(t1)

	ss := strings.Split(src, `\u`)
	dd := strings.Split(d1, `\u`)

	t.Error(ss[0], dd, d1, src)

}

func u2s1(form string) (to string, err error) {
	ss := strings.Split(form, "\\u")
	fmt.Println(ss)
	for _, s := range ss {
		if len(s) < 1 {
			continue
		}

		tmp, err := strconv.ParseInt(s, 16, 32)
		if nil != err {
			return "", err
		}

		to += fmt.Sprintf("%c", tmp)
	}
	return
}

func u2s(form string) (to string, err error) {
	bs, err := hex.DecodeString(strings.Replace(form, `\u`, ``, -1))
	if err != nil {
		return
	}
	for i, bl, br, r := 0, len(bs), bytes.NewReader(bs), uint16(0); i < bl; i += 2 {
		binary.Read(br, binary.BigEndian, &r)
		to += string(r)
	}
	return
}

// func TestGjson(t *testing.T){
// 	src := "[["ab"]]"
// }
