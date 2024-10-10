import assert from "node:assert";
import { existsSync, readFileSync, readdirSync } from "node:fs";
import { join } from "node:path";
import type { TestContext } from "node:test";

const loadSample = (day: number, path: string) => {
  const path_ = join(`../samples/day${day}`, path);
  return existsSync(path_)
    ? readFileSync(path_).toString().replaceAll("\r\n", "\n").trim()
    : "";
};

const loadCases = (day: number, part: number) => {
  const outputs = readdirSync(`../samples/day${day}`)
    .filter((f) => f.endsWith(".output.txt") && f.startsWith(`p${part}`))
    .map((f) => loadSample(day, f));

  return outputs.map((output, idx) => {
    const input =
      loadSample(day, `p${part}.s${idx + 1}.input.txt`) ||
      loadSample(day, `s${idx + 1}.input.txt`);
    return [idx + 1, input, output] as const;
  });
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
