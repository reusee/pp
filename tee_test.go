package pp

import "testing"

func TestTee(t *testing.T) {
	var src IntSrc
	n := 0
	src = func() (*int, IntSrc, error) {
		if n >= 10 {
			return nil, nil, nil
		}
		n++
		return PtrOf(n), src, nil
	}

	collect := func(ints *[]int) IntSink {
		var sink IntSink
		sink = func(v *int) (IntSink, error) {
			if v == nil {
				return nil, nil
			}
			*ints = append(*ints, *v)
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
		Discard[int, IntSink],
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
