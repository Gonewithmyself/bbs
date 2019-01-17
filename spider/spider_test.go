package spider

import (
	"fmt"
	"strconv"
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
			if got := extend(tt.args); got.St != tt.want {
				m := map[int]int{}
				idx := 0
				ss := got.St
				cnt := 0
				for i := range ss {
					if ss[i] == '[' {
						idx++
						m[idx] = i
					} else if ss[i] == ']' {
						switch idx {

						case 2:
							if cnt < 5 {
								println("##", ss[m[idx]:i+1])
							}

							cnt++
						case 4:
						}
						idx--
					}

				}

				t1 := "他们互相朗诵诗"
				t2 := "《柯林斯高阶英汉双解学习词典》"
				d1 := strconv.QuoteToASCII(t1)
				d2 := strconv.QuoteToASCII(t2)
				fmt.Println("xx", string(d1))
				fmt.Println("xx", string(d2))
				t.Errorf("extend() = %v, want %v", got.St[0], tt.want)
			}
		})
	}
}
