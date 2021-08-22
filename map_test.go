package pp

import "testing"

func TestMap(t *testing.T) {
	var src Src[int]
	n := 0
	src = func() (*int, Src[int], error) {
		if n == 5 {
			return nil, nil, nil
		}
		n++
		return PtrOf(n), src, nil
	}

	var ints []int
	var sink Sink[int]
	sink = func(value *int) (Sink[int], error) {
		if value == nil {
			return nil, nil
		}
		ints = append(ints, *value)
		return sink, nil
	}

	var ints2 []int
	if err := Copy(
		MapSrc(src, func(v int) int {
			ints2 = append(ints2, v)
			return v
		}, nil),
		MapSink(sink, func(v int) int {
			return v * 2
		}),
	); err != nil {
		t.Fatal(err)
	}

	if len(ints) != 5 {
		t.Fatal()
	}
	if len(ints2) != 5 {
		t.Fatal()
	}

	if ints2[2] != 3 {
		t.Fatalf("got %d", ints2[2])
	}
	if ints[2] != 6 {
		t.Fatalf("got %d", ints[2])
	}

}
