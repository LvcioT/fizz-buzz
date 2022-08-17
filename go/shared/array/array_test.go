package array_test

import (
	"fizz-buzz/shared/array"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	// integer type source and dest
	assert.Equal(
		t,
		[]int{2, 4, 6, 8, 10},
		array.Map(
			[]int{1, 2, 3, 4, 5},
			func(value int) int { return value * 2 },
		),
	)

	// string as source, int as dest
	assert.Equal(
		t,
		[]int{6, 6, 5, 5},
		array.Map(
			[]string{"Donald", "Mickey", "Goofy", "Pluto"},
			func(value string) int { return len(value) },
		),
	)
}

func TestEach(t *testing.T) {
	var whole string
	var sum int

	array.Each([]rune{'a', 'b', 'c'}, func(i int, value rune) {
		sum += i
		whole += string(value)
	})

	assert.Equal(t, "abc", whole)
	assert.Equal(t, 3, sum)
}

func TestFilter(t *testing.T) {
	assert.Equal(
		t,
		[]int{3, 6},
		array.Filter(
			[]int{1, 2, 3, 4, 5, 6},
			func(value int) bool { return value%3 == 0 },
		),
	)
}

func TestFirst(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}

	var result int
	var err error

	result, err = array.First(data, func(value int) bool { return value%3 == 0 })

	assert.Equal(t, 3, result)
	assert.Nil(t, err)

	result, err = array.First(data, func(value int) bool { return value%7 == 0 })

	assert.Equal(t, 0, result)
	assert.NotNil(t, err)
}

func TestIn(t *testing.T) {
	assert.True(
		t,
		array.In(
			[]int{1, 2, 3, 4, 5, 6},
			3,
		),
	)

	assert.False(
		t,
		array.In(
			[]int{1, 2, 3, 4, 5, 6},
			9,
		),
	)
}
