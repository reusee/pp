package pp

import "testing"

func BenchmarkCopy(b *testing.B) {

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

	var discard IntSink
	discard = func(v *int) (IntSink, error) {
		if v == nil {
			return nil, nil
		}
		return discard, nil
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := Copy(
			emit(128),
			discard,
		); err != nil {
			b.Fatal(err)
		}
	}
}
