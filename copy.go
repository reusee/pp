package pp

func Copy(src Src, sinks ...Sink) error {
	return CopyWithMutate(nil, src, sinks...)
}

type Mutate = func(func(Src) Src)

func CopyWithMutate(
	mutatePtr *Mutate,
	src Src,
	sinks ...Sink,
) error {

	if mutatePtr != nil {
		*mutatePtr = func(fn func(Src) Src) {
			cur := src
			src = fn(cur)
		}
	}

	for {
		if len(sinks) == 0 {
			break
		}

		value, err := src.Next()
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
