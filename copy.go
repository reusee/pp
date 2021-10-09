package pp

func Copy[
	Src interface {
		~func() (*T, Src, error)
	},
	Sink interface {
		~func(*T) (Sink, error)
	},
	T any,
](src Src, sinks ...Sink) error {
	for {
		if len(sinks) == 0 {
			break
		}

		value, err := Get[Src](&src)
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
