package pp

import "testing"

func BenchmarkCopy(b *testing.B) {

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

	var discard Sink
	discard = func(v any) (Sink, error) {
		if v == nil {
			return nil, nil
		}
		return discard, nil
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := Copy(
			emit(3000*10000),
			discard,
		); err != nil {
			b.Fatal(err)
		}
	}
}
