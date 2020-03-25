package pp

func Copy(src Src, sinks ...Sink) error {
	var err error
	for {

		var value any
		for value == nil {
			if src != nil {
				value, src, err = src()
				if err != nil {
					return err
				}
			} else {
				break
			}
		}

		if len(sinks) > 0 {
			for i := 0; i < len(sinks); {
				sink := sinks[i]
				if sink == nil {
					sinks[i] = sinks[len(sinks)-1]
					sinks = sinks[:len(sinks)-1]
					continue
				}
				sink, err = sink(value)
				if err != nil {
					return err
				}
				if sink == nil {
					sinks[i] = sinks[len(sinks)-1]
					sinks = sinks[:len(sinks)-1]
					continue
				}
				sinks[i] = sink
				i++
			}
		} else {
			break
		}

		if len(sinks) == 0 && src == nil {
			break
		}

	}
	return nil
}
