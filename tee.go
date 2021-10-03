package pp

func Tee(src Src, sinks ...Sink) Src {
	return TeeSrc(src, sinks, nil)
}

func TeeSrc(
	src Src,
	sinks []Sink,
	cont Src,
) Src {
	var ret Src
	ret = func() (any, Src, error) {
		value, err := src.Next()
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
