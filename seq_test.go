package pp

import "testing"

func TestSeq(t *testing.T) {
	var n int
	if err := Copy(
		Seq[int](1, 2, 3),
		CountSink(&n),
	); err != nil {
		t.Fatal(err)
	}
	if n != 3 {
		t.Fatalf("got %d", n)
	}
}
