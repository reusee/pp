package pp

import "testing"

func TestCapSrc(t *testing.T) {
	var values Values[int]
	if err := Copy(
		CapSrc(
			Seq(1, 2, 3, 4, 5),
			2,
			nil,
		),
		CollectValues(&values),
	); err != nil {
		t.Fatal(err)
	}
	if len(values) != 2 {
		t.Fatal()
	}
	if values[0] != 1 {
		t.Fatal()
	}
	if values[1] != 2 {
		t.Fatal()
	}

	values = values[:0]
	if err := Copy(
		CapSrc(
			Seq(1),
			2,
			nil,
		),
		CollectValues(&values),
	); err != nil {
		t.Fatal(err)
	}
	if len(values) != 1 {
		t.Fatal()
	}
	if values[0] != 1 {
		t.Fatal()
	}
}
