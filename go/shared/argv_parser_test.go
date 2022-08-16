package shared_test

import (
	"fizz-buzz/shared"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var defaultResult = shared.ArgsType{
	N_From: 1,
	N_To:   100,
	Print:  "series",
}

func TestNoArgumnets(t *testing.T) {
	os.Args = []string{"filename"}

	assert.Equal(t, defaultResult, shared.Parser())
}

func TestExtraArguments(t *testing.T) {
	os.Args = []string{"filename", "argument=value"}

	assert.Equal(t, defaultResult, shared.Parser())
}

func TestExpectedIntAruments(t *testing.T) {
	os.Args = []string{"filename", "nFrom=-5", "nTo=5"}

	expected := defaultResult

	expected.N_From = -5
	expected.N_To = 5

	assert.Equal(t, expected, shared.Parser())
}

func TestDefaultPrintArument(t *testing.T) {
	os.Args = []string{"filename", "print=everyelse"}

	assert.Equal(t, defaultResult, shared.Parser())
}

func TestExpectedPrintArument(t *testing.T) {
	os.Args = []string{"filename", "print=end"}

	expected := defaultResult

	expected.Print = "end"

	assert.Equal(t, expected, shared.Parser())
}
