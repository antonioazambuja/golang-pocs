package main

import (
	"testing"

	"github.com/antonioazambuja/golang-pocs/unit_test/calc"
	"github.com/stretchr/testify/assert"
)

func TestSumOperation(test *testing.T) {
	assert := assert.New(test)
	assert.Equal(float64(2), calc.Sum(1, 1))
}

func TestSubOperation(test *testing.T) {
	assert := assert.New(test)
	assert.Equal(float64(1), calc.Sub(2, 1))
}

func TestMultOperation(test *testing.T) {
	assert := assert.New(test)
	assert.Equal(float64(2), calc.Mult(2, 1))
}

func TestDivOperation(test *testing.T) {
	assert := assert.New(test)
	assert.Equal(float64(2), calc.Div(18, 9))
}
