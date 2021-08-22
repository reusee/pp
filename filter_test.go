package pp

import "testing"

func TestFilterSrc(t *testing.T) {
	var src Src[int]
	n := 0
	src = func() (*int, Src[int], error) {
		if n >= 10 {
			return nil, nil, nil
		}
		n++
		return PtrOf(n), src, nil
	}

	collect := func(ints *[]int) Sink[int] {
		var sink Sink[int]
		sink = func(v *int) (Sink[int], error) {
			if v == nil {
				return nil, nil
			}
			*ints = append(*ints, *v)
			return sink, nil
		}
		return sink
	}

	var ints []int
	if err := Copy(
		FilterSrc(
			src,
			func(v int) bool {
				return v%2 == 0
			},
			nil,
		),
		collect(&ints),
	); err != nil {
		t.Fatal(err)
	}

	if len(ints) != 5 {
		t.Fatal()
	}
	if ints[0] != 2 {
		t.Fatal()
	}
}

func TestFilterSink(t *testing.T) {
	var src Src[int]
	n := 0
	src = func() (*int, Src[int], error) {
		if n >= 10 {
			return nil, nil, nil
		}
		n++
		return PtrOf(n), src, nil
	}

	collect := func(ints *[]int) Sink[int] {
		var sink Sink[int]
		sink = func(v *int) (Sink[int], error) {
			if v == nil {
				return nil, nil
			}
			*ints = append(*ints, *v)
			return sink, nil
		}
		return sink
	}

	var even, odd []int
	if err := Copy(
		Tee(
			src,
			FilterSink(
				collect(&even),
				func(v int) bool {
					return v%2 == 0
				},
			),
			FilterSink(
				collect(&odd),
				func(v int) bool {
					return v%2 != 0
				},
			),
		),
		Discard[int],
	); err != nil {
		t.Fatal(err)
	}

	if len(even) != 5 {
		t.Fatal()
	}
	if even[0] != 2 {
		t.Fatal()
	}
	if len(odd) != 5 {
		t.Fatal()
	}
	if odd[0] != 1 {
		t.Fatal()
	}
}
