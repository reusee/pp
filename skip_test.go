package pp

import "testing"

func TestSkip(t *testing.T) {
	var values Values
	if err := Copy(
		SkipSrc(
			Seq(nil, 1, 2, nil, 3, 4, 5),
			1,
			nil,
		),
		CollectValues(&values),
	); err != nil {
		t.Fatal(err)
	}
	if len(values) != 4 {
		t.Fatal()
	}
	if values[0] != 2 {
		t.Fatal()
	}
}
