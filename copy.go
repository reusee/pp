package pp

func Copy[
	T any,
	Src interface {
		~func() (*T, Src, error)
	},
](src Src, sinks ...Sink[T]) error {
	for {
		if len(sinks) == 0 {
			break
		}

		value, err := Get[T, Src](&src)
		if err != nil {
			return err
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
