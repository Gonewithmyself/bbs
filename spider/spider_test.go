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
		want [][2]string
	}{
		{"xx", "recite", [][2]string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extend(tt.args); len(got.St) != -1 {
				t.Errorf("extend() = %v, want %v", got, tt.want)
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
