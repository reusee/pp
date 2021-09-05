package pp

import "testing"

func TestSeq(t *testing.T) {
	var n int
	if err := Copy(
		Seq[int, IntSrc](1, 2, 3),
		CountSink[int, IntSink](&n),
	); err != nil {
		t.Fatal(err)
	}
	if n != 3 {
		t.Fatalf("got %d", n)
	}

	if err := Copy[int, IntSrc, IntSink](
		Seq[int, IntSrc](1, 2, 3),
		Discard[int, IntSink],
	); err != nil {
		t.Fatal(err)
	}
}
