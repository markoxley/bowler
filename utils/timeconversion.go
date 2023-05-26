package utils

import (
	"fmt"
	"time"
)

func TimeToSQL(t time.Time) string {
	var y, m, d, h, mn, s, ns int
	y = t.Year()
	m = int(t.Month())
	d = t.Day()
	h = t.Hour()
	mn = t.Minute()
	s = t.Second()
	ns = t.Nanosecond()
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d.%d", y, m, d, h, mn, s, ns)
}
