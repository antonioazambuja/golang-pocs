package test

type CalcFunctionsMock interface {
	SumExpected(num1, num2 float64) float64
	SubExpected(num1, num2 float64) float64
	MultExpected(num1, num2 float64) float64
	DivExpected(num1, num2 float64) float64
}

type CalcMock struct {
	CalcFunctionsMock
	num1, num2 float64
}

func NewCalcMock(num1, num2 float64) *CalcMock {
	return &CalcMock{
		num1: num1,
		num2: num2,
	}
}

func (calc CalcMock) SumExpected(expectedValue float64) float64 {
	return expectedValue
}

func (calc CalcMock) SubExpected(expectedValue float64) float64 {
	return expectedValue
}

func (calc CalcMock) MultExpected(expectedValue float64) float64 {
	return expectedValue
}

func (calc CalcMock) DivExpected(expectedValue float64) float64 {
	return expectedValue
}
