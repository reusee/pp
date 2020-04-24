package pp

import "testing"

func TestMap(t *testing.T) {
	var src Src
	n := 0
	src = func() (any, Src, error) {
		if n == 5 {
			return nil, nil, nil
		}
		n++
		return n, src, nil
	}

	var ints []int
	var sink Sink
	sink = func(value any) (Sink, error) {
		if value == nil {
			return nil, nil
		}
		ints = append(ints, value.(int))
		return sink, nil
	}

	var ints2 []int
	if err := Copy(
		MapSrc(src, func(v any) any {
			ints2 = append(ints2, v.(int))
			return v
		}, nil),
		MapSink(sink, func(v any) any {
			return v.(int) * 2
		}, nil),
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
