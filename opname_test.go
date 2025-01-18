package opname

import (
	"testing"
	"time"
)

func Test(t *testing.T) {
	cases := []struct {
		prefix string
		ti     time.Time
		pretty string

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
			n := format(tt.prefix, tt.ti, tt.pretty)
			actual := n
			if actual != tt.want {
				t.Fatalf("want: %s, actual: %s", tt.want, actual)
			}
		})
	}
}
