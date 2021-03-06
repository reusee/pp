package pp

// Src provides stream of values
//
// returning value may be nil, which does not indicate the end of stream
// returning nil Src indicates the end of stream
type Src func() (any, Src, error)

// Next returns the next non-null value, or returns nil if Src is nil
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
