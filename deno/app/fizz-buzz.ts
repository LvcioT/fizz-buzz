import type {FizzBuzzResult, FizzBuzzValue} from './fizz-buzz.d.ts'

export const series = (from: number, to: number): FizzBuzzResult[] => {
  const results: FizzBuzzResult[] = [];

  for (let i = from; i <= to; i++) {
    results.push({
      step: i,
      value: step(i),
    });
  }

  return results;
};

export const step = (num:number): FizzBuzzValue => {
  if (num % 3 == 0 && num % 5 == 0) {
    return "FizzBuzz";
  }

  if (num % 3 == 0) {
    return "Fizz";
  }

  if (num % 5 == 0) {
    return "Buzz";
  }

  return num;
};

export const echo = (values: FizzBuzzResult[]):void => {
  values.forEach(({ step, value }) => console.info(step, value));
};
