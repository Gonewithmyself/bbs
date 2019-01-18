package spider

import (
	"fmt"
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

						case 1:
							if cnt < 5 {
								println("##", m[idx], i, ss[m[idx] : i+1][:200])
							}

							cnt++
						case 4:
						}
						idx--
					}

				}
				t.Errorf("extend() = %v, want %v", got.St[0], tt.want)
			}
		})
	}
}

func Test_gjson(t *testing.T) {
	src := `{"st":[[[["wo"],["shi"]],[["ni"],["ge"]],"xinhua"], [[["wo"],["shi"]],[["ni"],["ge"]],"xinhua"]]}`
	fmt.Println(get(src, "st"))
	fmt.Println(get(src, "st.#.0.#"))

	t.Error("")
}
