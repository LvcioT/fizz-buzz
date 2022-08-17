package shared

import (
	"fizz-buzz/shared/array"
	"os"
	"strconv"
	"strings"
)

type PrintValue string

var PrintValues = []string{"series", "end"}

type ArgsType struct {
	N_From int
	N_To   int
	Print  string
}

// default parameters, kept global as helper, just in case
var defaultParameters = ArgsType{
	N_From: 1,
	N_To:   100,
	Print:  PrintValues[0],
}

func Parser() ArgsType {
	// maps parameters split by '='
	params := make(map[string][]string)

	for _, arg := range os.Args[1:] {
		split := strings.Split(arg, "=")

		params[split[0]] = split[1:]
	}

	// prepare results with default
	results := defaultParameters
	var value []string
	var ok bool

	value, ok = params["nFrom"]
	if ok && len(value) == 1 {
		arg, err := strconv.Atoi(value[0])
		if err == nil {
			results.N_From = arg
		}
	}

	value, ok = params["nTo"]
	if ok && len(value) == 1 {
		arg, err := strconv.Atoi(value[0])
		if err == nil {
			results.N_To = arg
		}
	}

	value, ok = params["print"]
	if ok && len(value) == 1 {
		arg := value[0]

		if array.In(PrintValues, arg) {

			results.Print = arg
		}
	}

	return results
}
