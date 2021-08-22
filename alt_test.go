package pp

import (
	"testing"
)

func TestAlt(t *testing.T) {
	var src Src[int]
	n := 0
	src = func() (*int, Src[int], error) {
		if n >= 10 {
			return nil, nil, nil
		}
		n++
		return PtrOf(n), src, nil
	}

	countTo := func(max int, n *int) Sink[int] {
		var sink Sink[int]
		sink = func(v *int) (Sink[int], error) {
			if v == nil {
				return nil, nil
			}
			if *v > max {
				return nil, nil
			}
			*n = *v
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
