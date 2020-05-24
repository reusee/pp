package pp

func Seq(values ...any) Src {
	var src Src
	src = func() (any, Src, error) {
		if len(values) == 0 {
			return nil, nil, nil
		}
		value := values[0]
		values = values[1:]
		return value, src, nil
	}
	return src
}
