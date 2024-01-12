package helper

import "testing"

func TestFilterSpecChars(t *testing.T) {
	cases := []struct {
		src  string
		want string
	}{
		{"", ""},
		{"h", "h"},
		{"zcg", "zcg"},
		{"AgF💙(休假中)", "AgF(休假中)"},
		{"张成钢", "张成钢"},
	}

	for _, c := range cases {
		got := FilterSpecChars(c.src)
		if got != c.want {
			t.Errorf("src=%s, want=%s, got=%s\n", c.src, c.want, got)
		}
	}
}

func TestGetMaskName(t *testing.T) {
	cases := []struct {
		src  string
		want string
	}{
		{"", ""},
		{"h", "h"},
		{"张成钢", "**钢"},
		{"成钢", "*钢"},
		{"h钢", "*钢"},
		{"张成钢h", "***h"},
	}

	for _, c := range cases {
		got := GetMaskName(c.src)
		if got != c.want {
			t.Errorf("src=%s, want=%s, got=%s\n", c.src, c.want, got)
		}
	}
}
