package math

type Math struct {
	A int
	B int
}

func (m Math) Soma() int {
	return m.A + m.B
}
