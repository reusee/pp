package pp

type Values []any

func (v Values) Iter(cont Src) Src {
	n := 0
	var src Src
	src = func() (any, Src, error) {
		if n >= len(v) {
			return nil, cont, nil
		}
		value := v[n]
		n++
		return value, src, nil
	}
	return src
}

func CollectValues(p *Values) Sink {
	var sink Sink
	sink = func(v any) (Sink, error) {
		if v == nil {
			return nil, nil
		}
		*p = append(*p, v)
		return sink, nil
	}
	return sink
}
