import assert from "node:assert";
import { readFileSync, readdirSync } from "node:fs";
import { join } from "node:path";
import type { TestContext } from "node:test";

const glob = (path: string, end: string) =>
  readdirSync(path)
    .filter((f) => f.endsWith(end))
    .map((f) => join(path, f));

const loadFile = (path: string) =>
  readFileSync(path).toString().replaceAll("\r\n", "\n");

const loadCases = (day: number, part: number): [number, string, string][] => {
  const outputs = glob(`../samples/day${day}`, `.${part}.output.txt`).map(
    loadFile,
  );

  return glob(`../samples/day${day}`, ".input.txt")
    .map(loadFile)
    .map((source, i) => [i + 1, source, outputs[i]] as const);
};

export const runTests = (
  t: TestContext,
  day: number,
  part: number,
  func: (input: string) => { toString(): string },
) => {
  for (const [i, source, output] of loadCases(day, part))
    t.test(`sample${i}`, () => {
      assert.equal(func(source).toString(), output);
    });
};
