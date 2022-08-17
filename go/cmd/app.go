package main

import (
	"fizz-buzz/internal/argv_parser"
	"fizz-buzz/internal/core"
)

func main() {
	args := argv_parser.Parser()

	values := core.Series(args.N_From, args.N_To)

	if args.Print == argv_parser.PrintValues[0] {
		core.Print(values)
	} else {
		core.Print(values[len(values)-1:])
	}
}
