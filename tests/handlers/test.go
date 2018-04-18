package handlers

type Test struct {
}

func (*Test) Version() string {
	return "1.0.0"
}
