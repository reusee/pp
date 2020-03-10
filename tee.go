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
		var value any
		var err error
		for src != nil {
			value, src, err = src()
			if err != nil {
				return nil, nil, err
			}
			if value != nil {
				break
			}
		}
		for i := 0; i < len(sinks); {
			sink, err := sinks[i](value)
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
