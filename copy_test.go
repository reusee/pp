package pp

import (
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
