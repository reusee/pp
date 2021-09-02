package pp

type Values[T any] []T

func IterValues[
	T any,
	Src interface {
		~func() (*T, Src, error)
	},
](v Values[T], cont Src) Src {
	n := 0
	var src Src
	src = func() (*T, Src, error) {
		if n >= len(v) {
			return nil, cont, nil
		}
		value := &v[n]
		n++
		return value, src, nil
	}
	return src
}

func CollectValues[
	T any,
	Sink interface {
		~func(*T) (Sink, error)
	},
](p *Values[T]) Sink {
	return Tap(func(v *T) error {
		*p = append(*p, *v)
		return nil
	})
}
