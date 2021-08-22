package pp

type Values[T any] []T

func (v Values[T]) Iter(cont Src[T]) Src[T] {
	n := 0
	var src Src[T]
	src = func() (*T, Src[T], error) {
		if n >= len(v) {
			return nil, cont, nil
		}
		value := &v[n]
		n++
		return value, src, nil
	}
	return src
}

func CollectValues[T any](p *Values[T]) Sink[T] {
	return Tap(func(v *T) error {
		*p = append(*p, *v)
		return nil
	})
}
