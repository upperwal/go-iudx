package rs

type QueryLatest bool

func NewQueryLatest() *QueryLatest {
	return new(QueryLatest)
}

func (ql QueryLatest) String() string {
	return "options:latest"
}
