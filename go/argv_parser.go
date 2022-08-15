package main

import (
	"fizz-buzz/shared"
	"os"
	"strconv"
	"strings"
)

type PrintValue string

var printValues = []string{"series", "end"}

type ArgsType struct {
	nFrom int
	nTo   int
	print string
}

// default parameters, kept global as helper, just in case
var defaultParameters = ArgsType{
	nFrom: 1,
	nTo:   100,
	print: printValues[0],
}

func parser() ArgsType {
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
		if err != nil {
			results.nFrom = arg
		}
	}

	value, ok = params["nTo"]
	if ok && len(value) == 1 {
		arg, err := strconv.Atoi(value[0])
		if err != nil {
			results.nTo = arg
		}
	}

	value, ok = params["print"]
	if ok && len(value) == 1 {
		arg := value[0]

		if shared.Array_in(printValues, arg) {

			results.print = arg
		}
	}

	return results
}
