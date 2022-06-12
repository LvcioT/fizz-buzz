import type { Parameters , PrintValue} from "./args-parser.d.ts";

const defaultParameters: Parameters = {
  nFrom: 1,
  nTo: 100,
  print: "series",
};

const isPrintValue = (value: string): value is PrintValue => ['series', 'end'].includes(value as PrintValue)

export const parser = (denoArgs: string[]): Parameters => {
  const results = { ...defaultParameters };

  denoArgs.forEach((item: string): void => {
    const [param, value] = item.split("=");

    switch (param) {
      case "nTo":
      case "nFrom":
        results[param] = parseInt(value);
        break;
      case "print":
        results[param] = isPrintValue( value)  ? value : 'series' 
        break;
    }
  });

  // sanitize

  return results;
};
