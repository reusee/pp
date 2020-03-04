package pp

import "testing"

func TestTee(t *testing.T) {
	var src Src
	n := 0
	src = func() (any, Src, error) {
		if n >= 10 {
			return nil, nil, nil
		}
		n++
		return n, src, nil
	}

	collect := func(ints *[]int) Sink {
		var sink Sink
		sink = func(v any) (Sink, error) {
			if v == nil {
				return nil, nil
			}
			*ints = append(*ints, v.(int))
			return sink, nil
		}
		return sink
	}

	var ints1, ints2 []int
	if err := Copy(
		Tee(
			src,
			collect(&ints1),
			collect(&ints2),
		),
		Discard,
	); err != nil {
		t.Fatal(err)
	}

	if len(ints1) != 10 {
		t.Fatal()
	}
	if len(ints2) != 10 {
		t.Fatal()
	}

}
