package pp

func Copy(src Src, sink Sink) error {
	var err error
	for {
		var value any
	step_src:
		if src != nil {
			value, src, err = src()
			if err != nil {
				return err
			}
			if value == nil {
				goto step_src
			}
		}
		if sink != nil {
			sink, err = sink(value)
			if err != nil {
				return err
			}
		}
		if sink == nil && src == nil {
			break
		}
	}
	return nil
}
