import { parser } from "./args-parser.ts";
import { Parameters } from "./args-parser.d.ts";
import { assertEquals } from "./deps.dev.ts";

const defaultParameters: Parameters = {
  nFrom: 1,
  nTo: 100,
  print: "series",
};

Deno.test("no arguments gives default", () => {
  const args = parser([]);

  assertEquals(defaultParameters, args);
});

Deno.test("extra argument gives default", () => {
  const args = parser(["dummy=argument"]);

  assertEquals(defaultParameters, args);
});

Deno.test("argument gives effect", () => {
  const argv = parser(["nFrom=-5"]);

  const expected = { ...defaultParameters };
  expected.nFrom = -5;

  assertEquals(expected, argv);
});

Deno.test("print extra option argument gives default", () => {
  const argv = parser(["print=everything"]);

  assertEquals(defaultParameters, argv);
});

Deno.test("print end argument gives effect", () => {
  const argv = parser(["print=end"]);

  const expected = { ...defaultParameters };
  expected.print = "end";

  assertEquals(expected, argv);
});
