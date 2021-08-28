package pp

func Tee[T any](src Src[T], sinks ...Sink[T]) Src[T] {
	return TeeSrc(src, sinks, nil)
}

func TeeSrc[T any](
	src Src[T],
	sinks []Sink[T],
	cont Src[T],
) Src[T] {
	var ret Src[T]
	ret = func() (*T, Src[T], error) {
		value, err := Get(&src)
		if err != nil {
			return nil, nil, err
		}
		for i := 0; i < len(sinks); {
			sink := sinks[i]
			if sink == nil {
				sinks[i] = sinks[len(sinks)-1]
				sinks = sinks[:len(sinks)-1]
				continue
			}
			sink, err = sink(value)
			if err != nil {
				return nil, nil, err
			}
			if sink == nil {
				sinks[i] = sinks[len(sinks)-1]
				sinks = sinks[:len(sinks)-1]
			} else {
				sinks[i] = sink
				i++
			}
		}
		if value == nil && len(sinks) == 0 {
			return nil, cont, nil
		}
		return value, ret, nil
	}
	return ret
}
