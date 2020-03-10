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
	if err := Copy(
		MapSrc(src, func(v any) any {
			ints = append(ints, v.(int))
			return v
		}, nil),
		Discard,
	); err != nil {
		t.Fatal(err)
	}

	if len(ints) != 5 {
		t.Fatal()
	}
}
