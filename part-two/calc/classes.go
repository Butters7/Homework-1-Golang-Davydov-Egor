package calc

type iCalculatable interface {
	Calculate() float64
}

type plus struct {
	leftValue  iCalculatable
	rightValue iCalculatable
}

type minus struct {
	leftValue  iCalculatable
	rightValue iCalculatable
}

type multiply struct {
	leftValue  iCalculatable
	rightValue iCalculatable
}

type division struct {
	leftValue  iCalculatable
	rightValue iCalculatable
}

type value struct {
	value float64
}

func (p *plus) Calculate() float64 {
	return p.leftValue.Calculate() + p.rightValue.Calculate()
}

func (m *minus) Calculate() float64 {
	return m.leftValue.Calculate() - m.rightValue.Calculate()
}

func (m *multiply) Calculate() float64 {
	return m.leftValue.Calculate() * m.rightValue.Calculate()
}

func (d *division) Calculate() float64 {
	return d.leftValue.Calculate() / d.rightValue.Calculate()
}

func (v *value) Calculate() float64 {
	return v.value
}
