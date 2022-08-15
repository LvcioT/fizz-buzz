package shared

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var defaultResult = ArgsType{
	N_From: 1,
	N_To:   100,
	Print:  "series",
}

func TestNoArgumnets(t *testing.T) {
	assert.Equal(t, defaultResult, Parser())
}

func TestExtraArguments(t *testing.T) {
	os.Args = append(os.Args, "argument=value")

	assert.Equal(t, defaultResult, Parser())
}

func TestExpectedIntAruments(t *testing.T) {
	os.Args = append(os.Args, "nFrom=-5")
	os.Args = append(os.Args, "nTo=5")

	expected := defaultResult

	expected.N_From = -5
	expected.N_To = 5

	assert.Equal(t, expected, Parser())
}

/*

test('print extra option argument gives default', ()=> {
    process.argv = ['dummy1', 'dummy2', 'print=everything']

    const argv = parser()

    expect(argv).toStrictEqual({
        nFrom: -5,
        nTo: 100,
        print: 'series'
    })
})

test('print end argument gives effect', ()=> {
    process.argv = ['dummy1', 'dummy2', 'ptint=end']

    const argv = parser()

    expect(argv).toStrictEqual({
        nFrom: -5,
        nTo: 100,
        print: 'series'
    })
})

*/
