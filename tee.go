package pp

func Tee[
	Src interface {
		~func() (*T, Src, error)
	},
	Sink interface {
		~func(*T) (Sink, error)
	},
	T any,
](src Src, sinks ...Sink) Src {
	return TeeSrc[Src, Sink](src, sinks, nil)
}

func TeeSrc[
	Src interface {
		~func() (*T, Src, error)
	},
	Sink interface {
		~func(*T) (Sink, error)
	},
	T any,
](
	src Src,
	sinks []Sink,
	cont Src,
) Src {
	var ret Src
	ret = func() (*T, Src, error) {
		value, err := Get[Src](&src)
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
