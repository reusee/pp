package pp

import (
	"testing"
)

func TestAlt(t *testing.T) {
	var src Src
	n := 0
	src = func() (any, Src, error) {
		if n >= 10 {
			return nil, nil, nil
		}
		n++
		return n, src, nil
	}

	countTo := func(max int, n *int) Sink {
		var sink Sink
		sink = func(v any) (Sink, error) {
			if v == nil {
				return nil, nil
			}
			if v.(int) > max {
				return nil, nil
			}
			*n = v.(int)
			return sink, nil
		}
		return sink
	}

	var i int
	if err := Copy(
		src,
		Alt(
			countTo(5, &i),
			countTo(3, &i),
		),
	); err != nil {
		t.Fatal(err)
	}

	if i != 3 {
		t.Fatalf("got %d", i)
	}

}
