package pp

import "testing"

func TestMapFilterSrc(t *testing.T) {
	var values Values[int]
	if err := Copy(
		MapFilterSrc(
			Seq[int, IntSrc](1, 2, 3),
			func(i int) *int {
				if i == 2 {
					return nil
				}
				return &i
			},
			nil,
		),
		CollectValues[int, IntSink](&values),
	); err != nil {
		t.Fatal(err)
	}
	if len(values) != 2 {
		t.Fatal()
	}
	if values[0] != 1 {
		t.Fatal()
	}
	if values[1] != 3 {
		t.Fatal()
	}
}

func TestMapFilterSink(t *testing.T) {
	var values Values[int]
	if err := Copy(
		Seq[int, IntSrc](1, 2, 3),
		MapFilterSink(
			CollectValues[int, IntSink](&values),
			func(i int) *int {
				if i == 2 {
					return nil
				}
				return &i
			},
		),
	); err != nil {
		t.Fatal(err)
	}
	if len(values) != 2 {
		t.Fatal()
	}
	if values[0] != 1 {
		t.Fatal()
	}
	if values[1] != 3 {
		t.Fatal()
	}
}
