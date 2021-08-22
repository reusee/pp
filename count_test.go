package pp

import "testing"

func TestCount(t *testing.T) {
	var src Src[int]
	n := 0
	src = func() (*int, Src[int], error) {
		if n >= 10 {
			return nil, nil, nil
		}
		n++
		return PtrOf(n), src, nil
	}

	var c int
	var c2 int
	if err := Copy(
		CountSrc(&c2, src, nil),
		CountSink[int](&c),
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
