package calc

type CalcFunctions interface {
	Sum(num1, num2 float64) float64
	Sub(num1, num2 float64) float64
	Mult(num1, num2 float64) float64
	Div(num1, num2 float64) float64
}

type Calc struct {
	CalcFunctions
	num1, num2 float64
}

func NewCalc(num1, num2 float64) *Calc {
	return &Calc{
		num1: num1,
		num2: num2,
	}
}

func (calc Calc) Sum() float64 {
	return calc.num1 + calc.num2
}

func (calc Calc) Sub() float64 {
	return calc.num1 - calc.num2
}

func (calc Calc) Mult() float64 {
	return calc.num1 * calc.num2
}

func (calc Calc) Div() float64 {
	return calc.num1 / calc.num2
}
