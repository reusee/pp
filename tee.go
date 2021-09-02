package pp

func Tee[
	T any,
	Src interface {
		~func() (*T, Src, error)
	},
](src Src, sinks ...Sink[T]) Src {
	return TeeSrc(src, sinks, nil)
}

func TeeSrc[
	T any,
	Src interface {
		~func() (*T, Src, error)
	},
](
	src Src,
	sinks []Sink[T],
	cont Src,
) Src {
	var ret Src
	ret = func() (*T, Src, error) {
		value, err := Get[T, Src](&src)
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
