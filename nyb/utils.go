package nyb

import (
	"strings"
	"time"

	"github.com/hako/durafmt"
)

func removeMilliseconds(dur *durafmt.Durafmt) string {
	arr := strings.Split(dur.String(), " ")
	if len(arr) < 3 {
		return dur.String()
	}
	return strings.Join(arr[:len(arr)-2], " ")
}

func getOffset(target time.Time, zone *time.Location) int {
	_, offset := time.Date(target.Year(), target.Month(), target.Day(),
		target.Hour(), target.Minute(), target.Second(),
		target.Nanosecond(), zone).Zone()
	return offset
}

func pingpong(c chan bool) {
	select {
	case c <- true:
	default:
		return
	}
}

func changeNick(n string) string {
	if len(n) < 16 {
		n += "_"
		return n
	}
	n = strings.TrimRight(n, "_")
	if len(n) > 12 {
		n = n[:12] + "_"
	}
	return n
}

var timeNow = func() time.Time {
	return time.Now()
}
