package pp

import "testing"

func TestCount(t *testing.T) {
	var src Src
	n := 0
	src = func() (any, Src, error) {
		if n >= 10 {
			return nil, nil, nil
		}
		n++
		return n, src, nil
	}

	var c int
	var c2 int
	if err := Copy(
		CountSrc(&c2, src, nil),
		CountSink(&c),
	); err != nil {
		t.Fatal(err)
	}

	if c != 10 {
		t.Fatalf("got %d", c)
	}
	if c2 != 10 {
		t.Fatalf("got %d", c)
	}
}
