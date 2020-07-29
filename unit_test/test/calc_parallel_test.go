package main

import (
	"testing"

	"github.com/antonioazambuja/golang-pocs/unit_test/calc"

	"github.com/stretchr/testify/assert"
)

func TestParallelSumOperation(test *testing.T) {
	test.Parallel()
	assert := assert.New(test)
	assert.Equal(float64(2), calc.Sum(1, 1))
}

func TestParallelSubOperation(test *testing.T) {
	test.Parallel()
	assert := assert.New(test)
	assert.Equal(float64(1), calc.Sub(2, 1))
}

func TestParallelMultOperation(test *testing.T) {
	test.Parallel()
	assert := assert.New(test)
	assert.Equal(float64(2), calc.Mult(2, 1))
}

func TestParallelDivOperation(test *testing.T) {
	test.Parallel()
	assert := assert.New(test)
	assert.Equal(float64(2), calc.Div(18, 9))
}
