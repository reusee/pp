package pp

import (
	"strconv"
	"testing"
)

func TestCopy(t *testing.T) {
	var src Src
	n := 0
	src = func() (any, Src, error) {
		if n >= 10 {
			return nil, nil, nil
		}
		n++
		return n, src, nil
	}
	var sink Sink
	ns := ""
	sink = func(n any) (Sink, error) {
		if n == nil {
			return nil, nil
		}
		ns += strconv.Itoa(n.(int))
		return sink, nil
	}
	if err := Copy(src, sink); err != nil {
		t.Fatal(err)
	}
	if ns != "12345678910" {
		t.Fatal()
	}
}
