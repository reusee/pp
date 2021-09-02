package pp

type IntSrc func() (*int, IntSrc, error)
