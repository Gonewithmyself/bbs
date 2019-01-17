package spider

import "testing"

func Test_trans(t *testing.T) {

	tests := []struct {
		name string
		args string
		want string
	}{
		{"xx", "test", "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trans(tt.args); got != tt.want {
				t.Errorf("trans() = %v, want %v", got, tt.want)
			}
		})
	}
}
