package core

import (
	"fizz-buzz/shared"
	"fmt"
	"strconv"
)

type StepType struct {
	Step  int
	Value string
}

func Series(from, to int) []StepType {
	if from > to {
		return []StepType{}
	}

	length := to - from + 1

	results := make([]StepType, length)

	for i := 0; i < length; i++ {
		step := i + from
		results[i] = StepType{
			Step:  step,
			Value: Step(step),
		}
	}

	return results
}

func Step(num int) string {
	if num%3 == 0 && num%5 == 0 {
		return "FizzBuzz"
	}

	if num%3 == 0 {
		return "Fizz"
	}

	if num%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(num)
}

func Echo(values []StepType) {
	shared.Array_each(values, func(i int, item StepType) {
		fmt.Println(item)
	})
}
