import { series, step } from "./fizz-buzz.ts";
import { assertEquals } from "./deps.dev.ts";
import { FizzBuzzValue } from "./fizz-buzz.d.ts";

Deno.test("11 is 11", () => {
  assertEquals(11, step(11));
});

Deno.test("12 is Fizz", () => {
  assertEquals("Fizz" as FizzBuzzValue, step(12));
});

Deno.test("25 is FizzBuzz", () => {
  assertEquals("Buzz" as FizzBuzzValue, step(25));
});

Deno.test("30 is FizzBuzz", () => {
  assertEquals("FizzBuzz" as FizzBuzzValue, step(30));
});

Deno.test("from -10 to 10 are 21 items", () => {
  const values = series(-10, 10);

  assertEquals(values.length, 21);
});

Deno.test("from -10 to 0 first item is -10", () => {
  const values = series(-10, 0);

  assertEquals(values[0].step, -10);
});

Deno.test("from 0 to 10 last item is 10", () => {
  const values = series(0, 10);

  assertEquals(11, values.length);
  assertEquals(values.at(-1)?.step, 10);
});
