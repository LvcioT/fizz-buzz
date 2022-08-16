package main

import (
	"fizz-buzz/core"
	"fizz-buzz/shared"
)

func main() {
	args := shared.Parser()

	values := core.Series(args.N_From, args.N_To)

	if args.Print == shared.PrintValues[0] {
		core.Echo(values)
	} else {
		core.Echo(core.Series(args.N_To, args.N_To))
	}
}
