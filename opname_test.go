package opname

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	cases := []struct {
		prefix   string
		ti       time.Time
		nickname string

		want string
	}{
		{
			"usr",
			time.Date(2023, time.April, 01, 23, 45, 07, 12, time.Local),
			"soba",

			"usr20230401234507soba",
		},
	}

	for _, tt := range cases {
		t.Run(tt.want, func(t *testing.T) {
			n := format(tt.prefix, tt.ti, tt.nickname)
			actual := n
			if actual != tt.want {
				t.Fatalf("want: %s, actual: %s", tt.want, actual)
			}
		})
	}
}

func TestValidPrefix(t *testing.T) {
	cases := []struct {
		s    string
		want bool
	}{
		{"ok", true},
		{"safe", true},
		{"longer", false},
		{"", false},
		{"1ng", false},
		{"ok1", true},
		{"o1k", true},
	}

	for _, tt := range cases {
		t.Run(tt.s, func(t *testing.T) {
			if validPrefix(tt.s) != tt.want {
				t.Fail()
			}
		})
	}
}

func TestValidNickname(t *testing.T) {
	cases := []struct {
		s    string
		want bool
	}{
		{"ok", true},
		{"safe", true},
		{"longerlongerlongerlongerlonger", false},
		{"", false},
		{"1ng", true},
		{"ok1", false},
		{"o1k", true},
	}

	for _, tt := range cases {
		t.Run(tt.s, func(t *testing.T) {
			if validNichname(tt.s) != tt.want {
				t.Fail()
			}
		})
	}
}
