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
	return Tap(func(v any) {
		*p = append(*p, v)
	})
}
