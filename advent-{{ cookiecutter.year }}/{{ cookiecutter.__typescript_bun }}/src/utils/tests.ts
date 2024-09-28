import { expect, test } from "bun:test";
import { readFileSync, readdirSync } from "node:fs";
import { join } from "node:path";

const glob = (path: string, end: string) =>
  readdirSync(path)
    .filter((f) => f.endsWith(end))
    .map((f) => join(path, f));

const loadFile = (path: string) =>
  readFileSync(path).toString().replaceAll("\r\n", "\n");

const loadCases = (day: number, part: number) => {
  const outputs = glob(`../samples/day${day}`, `.${part}.output.txt`).map(
    loadFile,
  );

  return glob(`../samples/day${day}`, ".input.txt")
    .map(loadFile)
    .map((source, i) => [i + 1, source, outputs[i]] as const);
};

export const runTests = (
  day: number,
  part: number,
  func: (input: string) => { toString(): string },
) => {
  test.each(loadCases(day, part))("sample%d", (_, source, output) => {
    expect(func(source).toString()).toEqual(output);
  });
};
