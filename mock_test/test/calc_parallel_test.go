package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParallelSumOperation(test *testing.T) {
	test.Parallel()
	assert := assert.New(test)
	assert.Equal(float64(2), NewCalcMock(1, 1).SumExpected(float64(2)))
}

func TestParallelSubOperation(test *testing.T) {
	test.Parallel()
	assert := assert.New(test)
	assert.Equal(float64(1), NewCalcMock(2, 1).SubExpected(float64(1)))
}

func TestParallelMultOperation(test *testing.T) {
	test.Parallel()
	assert := assert.New(test)
	assert.Equal(float64(2), NewCalcMock(2, 1).MultExpected(float64(2)))
}

func TestParallelDivOperation(test *testing.T) {
	test.Parallel()
	assert := assert.New(test)
	assert.Equal(float64(2), NewCalcMock(18, 9).DivExpected(float64(2)))
}
