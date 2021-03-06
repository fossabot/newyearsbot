package nyb

import (
	"log"
	"strings"
	"testing"
)

type devNull struct {
}

func (d *devNull) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func init() {
	timeNow()
	log.SetOutput(&devNull{})
}

func TestChangeNick(t *testing.T) {
	ln := "aaaaaaaaaaaaaaaa"
	for i := 1; i <= 4; i++ {
		ln = changeNick(ln)
		if got := strings.Count(ln, "_"); got != i {
			t.Errorf("expecting %d _'s got %d, string: %s", i, got, ln)
		}
		if len(ln) != 12+i {
			t.Errorf("expecting length %d got %d, string: %s", 12+i, len(ln), ln)
		}
	}
	if ln := changeNick(ln); len(ln) != 12 {
		t.Errorf("expecting length 12 got %d, string: %s", len(ln), ln)
	}
	n := "a"
	for i := 1; i <= 15; i++ {
		n = changeNick(n)
		c := strings.Count(n, "_")
		if c != i {
			t.Errorf("expecting %d _'s got %d, string: %s", i, c, n)
		}
	}
	if n := changeNick(n); n != "a" {
		t.Error("nick wasn't 'a'")
	}
}
