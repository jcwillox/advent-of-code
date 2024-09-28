import { test } from "node:test";
import * as day from "./day{{ cookiecutter.day }}.ts";
import { runTests } from "./utils/tests.ts";

test("part1", (t) => {
  runTests(t, 1, 1, day.part1);
});

test("part2", (t) => {
  runTests(t, 1, 2, day.part2);
});
