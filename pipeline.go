package pp

type Src func() (any, Src, error)

type Sink func(any) (Sink, error)

func (s *Src) Next() (any, error) {
	var value any
	var err error
	for value == nil {
		if s != nil && *s != nil {
			value, *s, err = (*s)()
			if err != nil {
				return nil, err
			}
		} else {
			break
		}
	}
	return value, nil
}
