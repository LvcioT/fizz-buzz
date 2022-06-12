
export type PrintValue = "series" | "end";

export type Parameters = {
  [key: string]: number | PrintValue;
  nFrom: number;
  nTo: number;
  print: PrintValue;
};
