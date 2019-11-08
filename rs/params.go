package rs

type QueryLatest bool

func NewQueryLatest() *QueryLatest {
	return new(QueryLatest)
}

func (ql QueryLatest) Params() []string {
	return []string{"options:latest"}
}
