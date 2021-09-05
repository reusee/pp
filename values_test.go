package pp

import "testing"

func TestValues(t *testing.T) {
	values := Values[int]{1, 2, 3}
	var v2 Values[int]
	if err := Copy(
    IterValues[int, IntSrc](values, nil),
		CollectValues[int, IntSink](&v2),
	); err != nil {
		t.Fatal(err)
	}
	if len(v2) != 3 {
		t.Fatal()
	}
	if v2[0] != 1 {
		t.Fatal()
	}
	if v2[1] != 2 {
		t.Fatal()
	}
	if v2[2] != 3 {
		t.Fatal()
	}
}
