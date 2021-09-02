package pp

type IntSink func(*int) (IntSink, error)
