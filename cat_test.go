package pp

import "testing"

func TestCatSrc(t *testing.T) {
	seq := func(a, b int) Src {
		var src Src
		src = func() (any, Src, error) {
			if a == b {
				return nil, nil, nil
			}
			a++
			return a - 1, src, nil
		}
		return src
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

	var ints []int
	if err := Copy(
		CatSrc(
			seq(0, 1),
			seq(1, 3),
			seq(3, 6),
		),
		collect(&ints),
	); err != nil {
		t.Fatal(err)
	}
	if len(ints) != 6 {
		t.Fatal()
	}
	for i, n := range ints {
		if n != i {
			t.Fatal()
		}
	}

}
