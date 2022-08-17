package core_test

import (
	"fizz-buzz/internal/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStep(t *testing.T) {
	type Test_Case struct {
		input  int
		output string
	}

	cases := []Test_Case{
		{11, "11"},
		{12, "Fizz"},
		{25, "Buzz"},
		{30, "FizzBuzz"},
	}

	for _, test_case := range cases {
		assert.Equal(t, test_case.output, core.Step(test_case.input))
	}
}

func TestSeries(t *testing.T) {
	var values []core.StepType

	//  total elements
	values = core.Series(-10, 10)
	assert.Len(t, values, 21)

	// first item (even negative)
	values = core.Series(-10, 0)
	assert.Equal(t, -10, values[0].Step)

	// last item
	values = core.Series(0, 10)
	assert.Equal(t, 10, values[len(values)-1].Step)
}
