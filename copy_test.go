package pp

import (
	"strconv"
	"testing"
)

func TestCopy(t *testing.T) {
	var src IntSrc
	n := 0
	src = func() (*int, IntSrc, error) {
		if n >= 10 {
			return nil, nil, nil
		}
		n++
		return PtrOf(n), src, nil
	}
	var sink Sink[int]
	ns := ""
	sink = func(n *int) (Sink[int], error) {
		if n == nil {
			return nil, nil
		}
		ns += strconv.Itoa(*n)
		return sink, nil
	}
	if err := Copy(src, sink); err != nil {
		t.Fatal(err)
	}
	if ns != "12345678910" {
		t.Fatal()
	}
}

func TestCopyToMultipleSinks(t *testing.T) {
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

	emit := func(n int) IntSrc {
		var src IntSrc
		src = func() (*int, IntSrc, error) {
			if n == 0 {
				return nil, nil, nil
			}
			n--
			return PtrOf(n), src, nil
		}
		return src
	}

	var s1, s2, s3 []int
	if err := Copy(
		emit(10),
		collect(&s1),
		collect(&s2),
		collect(&s3),
	); err != nil {
		t.Fatal(err)
	}

	if len(s1) != 10 {
		t.Fatalf("got %d, %v", len(s1), s1)
	}
	if len(s2) != 10 {
		t.Fatalf("got %d, %v", len(s2), s2)
	}
	if len(s3) != 10 {
		t.Fatalf("got %d, %v", len(s3), s3)
	}
}

func TestCopyMultipleValues(t *testing.T) {
	var provide IntSrc
	ints := []int{1, 2, 3}
	provide = func() (*int, IntSrc, error) {
		if len(ints) == 0 {
			return nil, nil, nil
		}
		i := ints[0]
		ints = ints[1:]
		return PtrOf(i), provide, nil
	}
	consume := func(target *int, cont Sink[int]) Sink[int] {
		return func(value *int) (Sink[int], error) {
			*target = *value
			return cont, nil
		}
	}
	var a, b int
	if err := Copy(
		provide,
		consume(&a, consume(&b, nil)),
	); err != nil {
		t.Fatal(err)
	}
	if a != 1 {
		t.Fatal()
	}
	if b != 2 {
		t.Fatal()
	}
	if len(ints) != 1 {
		t.Fatal()
	}
	if ints[0] != 3 {
		t.Fatal()
	}
}
