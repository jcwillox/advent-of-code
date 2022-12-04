import { readFileSync } from "fs";
import fg from "fast-glob";

const loadCases = (day: number, part: number): [number, string, string][] => {
  const loadFile = (path: string) =>
    readFileSync(path).toString().replaceAll("\r\n", "\n");

  const outputs = fg
    .sync(`../samples/day${day}/*.${part}.output.txt`)
    .map(loadFile);

  return fg
    .sync(`../samples/day${day}/*.input.txt`)
    .map(loadFile)
    .map<[number, string, string]>((source, i) => {
      return [i + 1, source, outputs[i]];
    });
};

export const runTests = (
  day: number,
  part: number,
  func: (input: string) => { toString(): string }
) => {
  test.each(loadCases(day, part))("sample%d", (_, source, output) => {
    assert.equal(func(source).toString(), output);
  });
};
