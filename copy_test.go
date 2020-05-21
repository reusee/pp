package pp

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCopy(t *testing.T) {
	var src Src
	n := 0
	src = func() (any, Src, error) {
		if n >= 10 {
			return nil, nil, nil
		}
		n++
		return n, src, nil
	}
	var sink Sink
	ns := ""
	sink = func(n any) (Sink, error) {
		if n == nil {
			return nil, nil
		}
		ns += strconv.Itoa(n.(int))
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

	emit := func(n int) Src {
		var src Src
		src = func() (any, Src, error) {
			if n == 0 {
				return nil, nil, nil
			}
			n--
			return n, src, nil
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
	var provide Src
	ints := []int{1, 2, 3}
	provide = func() (any, Src, error) {
		if len(ints) == 0 {
			return nil, nil, nil
		}
		i := ints[0]
		ints = ints[1:]
		return i, provide, nil
	}
	consume := func(target *int, cont Sink) Sink {
		return func(value any) (Sink, error) {
			*target = value.(int)
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

func TestCopyWithMutate(t *testing.T) {
	var ones Src
	ones = func() (any, Src, error) {
		return 1, ones, nil
	}

	var mutate Mutate
	var insertTwos Sink
	var ints []int
	insertTwos = func(v any) (Sink, error) {
		i := v.(int)
		ints = append(ints, i)
		if i == 1 {
			// insert 42 after 1
			mutate(func(src Src) Src {
				return func() (any, Src, error) {
					return 42, src, nil
				}
			})
		}
		if len(ints) == 8 {
			return nil, nil
		}
		return insertTwos, nil
	}

	if err := CopyWithMutate(&mutate, ones, insertTwos); err != nil {
		t.Fatal(err)
	}

	if len(ints) != 8 {
		t.Fatalf("got %v", ints)
	}
	if fmt.Sprintf("%v", ints) != "[1 42 1 42 1 42 1 42]" {
		t.Fatalf("got %v", ints)
	}

}
