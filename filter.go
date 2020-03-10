package pp

func FilterSrc(
	src Src,
	predict func(any) bool,
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
		if src == nil && value == nil {
			return nil, cont, nil
		}
		if value != nil && !predict(value) {
			value = nil
		}
		return value, ret, nil
	}
	return ret
}

func FilterSink(
	sink Sink,
	predict func(any) bool,
) Sink {
	var ret Sink
	ret = func(value any) (Sink, error) {
		var err error
		if value == nil || predict(value) {
			if sink == nil {
				return nil, nil
			}
			sink, err = sink(value)
			if err != nil {
				return nil, err
			}
		}
		if sink == nil {
			return nil, nil
		}
		return ret, nil
	}
	return ret
}
