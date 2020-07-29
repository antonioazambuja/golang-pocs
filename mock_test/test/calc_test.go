package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOperation(test *testing.T) {
	assert := assert.New(test)
	assert.Equal(float64(2), NewCalcMock(1, 1).SumExpected(float64(2)))
}

func TestSubOperation(test *testing.T) {
	assert := assert.New(test)
	assert.Equal(float64(1), NewCalcMock(2, 1).SubExpected(float64(1)))
}

func TestMultOperation(test *testing.T) {
	assert := assert.New(test)
	assert.Equal(float64(2), NewCalcMock(2, 1).MultExpected(float64(2)))
}

func TestDivOperation(test *testing.T) {
	assert := assert.New(test)
	assert.Equal(float64(2), NewCalcMock(18, 9).DivExpected(float64(2)))
}
