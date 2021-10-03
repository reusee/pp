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

	collect := func(p *int) Sink {
		return func(value any) (Sink, error) {
			*p = value.(int)
			return nil, nil
		}
	}

	var a, b, c, d, e, f int
	if err := Copy(
		CatSrc(
			seq(0, 1),
			seq(1, 3),
			seq(3, 6),
		),
		CatSink(
			collect(&a),
			collect(&b),
			collect(&c),
			collect(&d),
			collect(&e),
			collect(&f),
		),
	); err != nil {
		t.Fatal(err)
	}
	if a != 0 {
		t.Fatal()
	}
	if b != 1 {
		t.Fatal()
	}
	if c != 2 {
		t.Fatal()
	}
	if d != 3 {
		t.Fatal()
	}
	if e != 4 {
		t.Fatal()
	}
	if f != 5 {
		t.Fatal()
	}

}
