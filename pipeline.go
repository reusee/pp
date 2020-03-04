package pp

type Src func() (any, Src, error)

type Sink func(any) (Sink, error)
