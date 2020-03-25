package pp

type Src func() (any, Src, error)

type Sink func(any) (Sink, error)

func (s *Src) Next() (value any, err error) {
	for value == nil {
		if s != nil && *s != nil {
			value, *s, err = (*s)()
		} else {
			break
		}
	}
	return
}
